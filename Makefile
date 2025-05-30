# /path/to/your/copa-wiz-repo/Makefile

CLI_BINARY=copa-wiz

.PHONY: all build clean release

all: build

build:
	go build -o dist/linux_amd64/release/$(CLI_BINARY) .

clean:
	rm -rf dist/

release: clean
	mkdir -p dist/linux_amd64/release
	mkdir -p dist/linux_arm64/release
	mkdir -p dist/darwin_amd64/release
	mkdir -p dist/darwin_arm64/release
	mkdir -p dist/windows_amd64/release

	echo "Building for Linux AMD64..."
	GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/release/$(CLI_BINARY) .
	tar -czf dist/$(CLI_BINARY)-linux-amd64.tar.gz -C dist/linux_amd64/release $(CLI_BINARY)

	echo "Building for Linux ARM64..."
	GOOS=linux GOARCH=arm64 go build -o dist/linux_arm64/release/$(CLI_BINARY) .
	tar -czf dist/$(CLI_BINARY)-linux-arm64.tar.gz -C dist/linux_arm64/release $(CLI_BINARY)

	echo "Building for macOS AMD64..."
	GOOS=darwin GOARCH=amd64 go build -o dist/darwin_amd64/release/$(CLI_BINARY) .
	tar -czf dist/$(CLI_BINARY)-darwin-amd64.tar.gz -C dist/darwin_amd64/release $(CLI_BINARY)

	echo "Building for macOS ARM64 (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build -o dist/darwin_arm64/release/$(CLI_BINARY) .
	tar -czf dist/$(CLI_BINARY)-darwin-arm64.tar.gz -C dist/darwin_arm64/release $(CLI_BINARY)

	echo "Building for Windows AMD64..."
	GOOS=windows GOARCH=amd64 go build -o dist/windows_amd64/release/$(CLI_BINARY).exe .
	zip -j dist/$(CLI_BINARY)-windows-amd64.zip dist/windows_amd64/release/$(CLI_BINARY).exe

	echo "Release binaries built in the 'dist' directory."