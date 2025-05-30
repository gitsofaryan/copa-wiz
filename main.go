// copa-wiz/main.go
// This is the main Go program for the copa-wiz plugin.
// It parses a Wiz JSON vulnerability report and converts it into
// Copacetic's v1alpha1.UpdateManifest format.
package main

import (
	"encoding/json"
	"errors" // Import errors package for custom errors
	"flag"
	"fmt"
	"os"

	v1alpha1 "github.com/project-copacetic/copacetic/pkg/types/v1alpha1"
)

// WizParser struct to encapsulate Wiz report parsing logic.
type WizParser struct{}

// NewWizParser creates and returns a new instance of WizParser.
func NewWizParser() *WizParser {
	return &WizParser{}
}

// wizReport defines the expected structure of the Wiz JSON report.
// This is a placeholder structure as noted in the original PR description,
// pending an actual sample Wiz report.
type wizReport struct {
	OS struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"os"`
	Vulnerabilities []struct {
		PackageName      string `json:"packageName"`
		InstalledVersion string `json:"installedVersion"`
		FixedVersion     string `json:"fixedVersion"`
		CVEID            string `json:"cveId"`
	} `json:"vulnerabilities"`
	// Add a field here if Wiz reports have a specific identifier,
	// similar to Grype's Descriptor.Name, to validate the report format.
	// Example: Descriptor struct { Name string `json:"name"` } `json:"descriptor"`
}

// parseWizReport reads a JSON file and unmarshals it into a wizReport struct.
func parseWizReport(filePath string) (*wizReport, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var report wizReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

// Parse method implements the logic to parse a Wiz report file
// and convert it into a v1alpha1.UpdateManifest.
func (wp *WizParser) Parse(filePath string) (*v1alpha1.UpdateManifest, error) {
	// Parse the Wiz scan results
	report, err := parseWizReport(filePath)
	if err != nil {
		return nil, err
	}

	// --- Placeholder for Wiz report format validation ---
	// Similar to how Grype checks `report.Descriptor.Name != "grype"`,
	// you would add a check here for Wiz specific fields to ensure
	// the report is in the expected format.
	// For example, if Wiz reports have a "scanner" field:
	// if report.Scanner.Name != "wiz" {
	//     return nil, errors.New("report format not supported by wiz: scanner name mismatch")
	// }
	// For now, we'll assume if it unmarshals, it's "valid enough" based on the placeholder struct.
	if report.OS.Name == "" && len(report.Vulnerabilities) == 0 {
		return nil, errors.New("wiz report appears empty or malformed (no OS info or vulnerabilities)")
	}
	// --- End Placeholder ---

	updates := v1alpha1.UpdateManifest{
		APIVersion: v1alpha1.APIVersion, // Use the APIVersion from the Copacetic types.
		Metadata: v1alpha1.Metadata{
			OS: v1alpha1.OS{
				Type:    report.OS.Name,    // OS name from Wiz report.
				Version: report.OS.Version, // OS version from Wiz report.
			},
			Config: v1alpha1.Config{
				Arch: "amd64", // Hardcoded architecture; update if Wiz report provides this dynamically.
				// If Wiz report provides architecture, it would be extracted here, e.g.:
				// Arch: report.Source.Target.(map[string]interface{})["architecture"].(string),
			},
		},
		Updates: make([]v1alpha1.UpdatePackage, 0), // Initialize as empty slice
	}

	// Iterate through the vulnerabilities in the Wiz report.
	// You can add filtering logic here similar to Grype's
	// (e.g., checking for OS-level packages or fixable vulnerabilities)
	// once the exact requirements for Wiz are known.
	for _, vuln := range report.Vulnerabilities {
		// Example: If you only want to include vulnerabilities that have a fixed version
		if vuln.FixedVersion != "" {
			updates.Updates = append(updates.Updates, v1alpha1.UpdatePackage{
				Name:             vuln.PackageName,
				InstalledVersion: vuln.InstalledVersion,
				FixedVersion:     vuln.FixedVersion,
				VulnerabilityID:  vuln.CVEID,
			})
		}
	}

	return &updates, nil
}

// main is the entry point of the copa-wiz plugin.
// It expects a single command-line argument: the path to the Wiz report file.
func main() {
	flag.Parse()
	// Check if exactly one argument (the report file path) is provided.
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: copa-wiz <report-file>")
		os.Exit(1)
	}

	reportFile := flag.Arg(0)
	parser := NewWizParser() // Create a new WizParser instance.

	// Parse the Wiz report using the Parser's Parse method.
	manifest, err := parser.Parse(reportFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing report: %v\n", err)
		os.Exit(1)
	}

	// Encode the generated UpdateManifest to JSON and print it to standard output.
	if err := json.NewEncoder(os.Stdout).Encode(manifest); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding manifest: %v\n", err)
		os.Exit(1)
	}
}
