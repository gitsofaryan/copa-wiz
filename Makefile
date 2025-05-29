(* copa-wiz/Makefile *)
(* This Makefile defines the build process for the copa-wiz plugin binary. *)
CLI_BINARY=copa-wiz

.PHONY: all build
all: build

build:
	go build -o dist/linux_amd64/release/$(CLI_BINARY) .
