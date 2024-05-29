VERSION := v0.1.1
INSTALL_DIR := ./bin
BINARY_NAME := wescale_wasm

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

DOWNLOAD_URL := https://github.com/wesql/wescale-wasm-plugin/releases/download/$(VERSION)/$(BINARY_NAME)_$(VERSION)_$(OS)_$(ARCH)

.PHONY: install-wescale-wasm
install-wescale-wasm:
	mkdir -p bin
	@echo "Downloading binary..."
	@curl -L -o $(BINARY_NAME) "$(DOWNLOAD_URL)"
	@if [ $$? -ne 0 ]; then \
		echo "Download failed, please check your network connection and URL correctness."; \
		exit 1; \
	fi
	@echo "Moving binary to $(INSTALL_DIR)..."
	@sudo mv $(BINARY_NAME) "$(INSTALL_DIR)"
	@echo "Setting executable permissions..."
	@sudo chmod +x "$(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "Installation completed. You can now use the $(BINARY_NAME) command."

.PHONY: uninstall-wescale-wasm
uninstall-wescale-wasm:
	@echo "Removing binary from $(INSTALL_DIR)..."
	@sudo rm -f "$(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "Uninstallation completed."

build-examples:
	mkdir -p bin
	# Iterate over all the examples and build them
	for example in $(shell ls ./examples); do \
		echo "Building example: $$example"; \
		tinygo build --no-debug -o ./bin/$$example.wasm -target=wasi -scheduler=none ./examples/$$example/main.go; \
	done

clean:
	rm -f ./bin/*
	rm -rf ./dist/*

# Default output filename
OUTPUT ?= guest.wasm

build:
	mkdir -p bin
	tinygo build --no-debug -o ./bin/$(OUTPUT) -target=wasi -scheduler=none ./main.go

install: build install-wescale-wasm
	./bin/wescale_wasm --command=install --wasm_file=./bin/$(OUTPUT)

uninstall:
	./bin/wescale_wasm --command=uninstall

