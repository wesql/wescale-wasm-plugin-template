VERSION := v0.1.2
INSTALL_DIR := ./bin
BINARY_NAME := wescale_wasm

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

DOWNLOAD_URL := https://github.com/wesql/wescale-wasm-plugin-sdk/releases/download/$(VERSION)/$(BINARY_NAME)_$(VERSION)_$(OS)_$(ARCH)

.PHONY: install-wescale-wasm
install-wescale-wasm:
	mkdir -p bin
	@if [ -f "$(INSTALL_DIR)/$(BINARY_NAME)" ]; then \
		echo "Binary already exists in $(INSTALL_DIR). Skipping installation."; \
	else \
		echo "Downloading binary..."; \
		curl -L -o $(BINARY_NAME) "$(DOWNLOAD_URL)"; \
		if [ $$? -ne 0 ]; then \
			echo "Download failed, please check your network connection and URL correctness."; \
			exit 1; \
		fi; \
		echo "Moving binary to $(INSTALL_DIR)..."; \
		mv $(BINARY_NAME) "$(INSTALL_DIR)"; \
		echo "Setting executable permissions..."; \
		chmod +x "$(INSTALL_DIR)/$(BINARY_NAME)"; \
		echo "Installation completed. You can now use the $(BINARY_NAME) command."; \
	fi

.PHONY: uninstall-wescale-wasm
uninstall-wescale-wasm:
	@echo "Removing binary from $(INSTALL_DIR)..."
	@rm -f "$(INSTALL_DIR)/$(BINARY_NAME)"
	@echo "Uninstallation completed."

.PHONY: build-examples
build-examples:
	mkdir -p bin
	# Iterate over all the examples and build them
	for example in $(shell ls ./examples); do \
		echo "Building example: $$example"; \
		cd ./examples/$$example && \
		tinygo build --no-debug -o ../../bin/$$example.wasm -target=wasi -scheduler=none ./main.go && \
		cd -; \
	done

MYSQL_HOST=127.0.0.1
MYSQL_PORT=15306
MYSQL_USER=root
MYSQL_PASSWORD=root

.PHONY: deploy-examples
deploy-examples:
	./bin/wescale_wasm --command=install --wasm_file=./bin/auditlog.wasm \
		--mysql_host=$(MYSQL_HOST) --mysql_port=$(MYSQL_PORT) --mysql_user=$(MYSQL_USER) --mysql_password=$(MYSQL_PASSWORD) \
		--filter_name=auditlog_wasm_filter --filter_desc='some kind of description' --filter_status=ACTIVE

	./bin/wescale_wasm --command=install --wasm_file=./bin/datamasking.wasm \
		--mysql_host=$(MYSQL_HOST) --mysql_port=$(MYSQL_PORT) --mysql_user=$(MYSQL_USER) --mysql_password=$(MYSQL_PASSWORD) \
		--filter_name=datamasking_wasm_filter --filter_desc='some kind of description' --filter_status=ACTIVE \
		--filter_plans="Select"

	./bin/wescale_wasm --command=install --wasm_file=./bin/interceptor.wasm \
		--mysql_host=$(MYSQL_HOST) --mysql_port=$(MYSQL_PORT) --mysql_user=$(MYSQL_USER) --mysql_password=$(MYSQL_PASSWORD) \
		--filter_name=interceptor_wasm_filter --filter_desc='some kind of description' --filter_status=ACTIVE \
		--filter_plans="Update,Delete"

undeploy-examples:
	./bin/wescale_wasm --command=uninstall --filter_name=auditlog_wasm_filter
	./bin/wescale_wasm --command=uninstall --filter_name=datamasking_wasm_filter
	./bin/wescale_wasm --command=uninstall --filter_name=interceptor_wasm_filter

.PHONY: clean
clean:
	rm -f ./bin/*

# Default output filename
WASM_FILE ?= my_plugin.wasm

.PHONY: build
build:
	mkdir -p bin
	tinygo build --no-debug -o ./bin/$(WASM_FILE) -target=wasi -scheduler=none ./main.go
