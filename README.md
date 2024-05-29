# WeScale-Wasm-Plugin-Template

## Pre-requisites
You need to have TinyGo to compile the code, you can install it by following the instructions [here](https://tinygo.org/getting-started/install/).

## Build

you should write your code in the `main.go` and then run the following command to build the plugin.
```bash
make build

# or if you want to specify the output file name
make build OUTPUT=my_plugin.wasm
```

If you want to build the example plugin, you can run the following command.
```bash
make build-examples
```

## Install
TODO

## Uninstall
TODO