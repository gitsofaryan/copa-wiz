# 🔌 Copa-Wiz

Plugin for [Copacetic](https://github.com/project-copacetic/copacetic) to support patching vulnerabilities identified by Wiz.

Learn more about Copacetic's scanner plugins [here](https://project-copacetic.github.io/copacetic/website/scanner-plugins).

## Installation

You can download the latest and previous versions of `copa-wiz` from the GitHub releases page (once available). Make sure to add it to your `PATH` environment variable.

Otherwise, install using the CLI:

```bash
# Install the binary (replace with actual release URL once available)
curl -sL [https://github.com/your-username/copa-wiz/releases/latest/download/copa-wiz](https://github.com/your-username/copa-wiz/releases/latest/download/copa-wiz) -o copa-wiz
```

## Example Usage
```

# generate a Wiz report (replace with actual Wiz CLI command)
# For now, assume a command like: wiz scan <image> -o json --file wiz_report.json
# You will need to use the actual Wiz CLI to generate this report.
wiz scan my-image -o json --file wiz_report.json

# test plugin with example config
copa-wiz wiz_report.json
# This will print the report in JSON format. Example (based on placeholder structure):
# {"apiVersion":"v1alpha1","metadata":{"os":{"type":"Debian","version":"11"},"config":{"arch":"amd64"}},"updates":[{"name":"openssl","installedVersion":"1.1.1n-0+deb11u5","fixedVersion":"1.1.1o-0+deb11u6","vulnerabilityID":"CVE-2022-0778"}]}

# run copa with the scanner plugin (copa-wiz) and the report file
copa patch -i $IMAGE -r wiz_report.json --scanner wiz
# run copa with the scanner plugin (copa-wiz) and the report file
copa patch -i $IMAGE -r wiz_report.json --scanner wiz
```
