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



