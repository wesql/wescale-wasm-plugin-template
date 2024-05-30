# WeScale-Wasm-Plugin-Template

## Pre-requisites
* You need to have TinyGo to compile the code, you can install it by following the instructions [here](https://tinygo.org/getting-started/install/).

* You will also need to have the `wescale_wasm` binary to deploy the plugin, you can download it by running the following command.
```bash
make install-wescale-wasm
```

## Build

you should write your code in the `main.go` and then run the following command to build the plugin.
```bash
make build

# or if you want to specify the output file name
make build WASM_FILE=my_plugin.wasm
```

If you want to build the example plugin, you can run the following command.
```bash
make build-examples
```

## Deploy
**How to deploy a wasm binary to WeScale to be a Filter:**
1. You need to use the `make build` command to build the wasm binary (if it's not already built).
2. You need to use the `make install-wescale-wasm` command to download the `wescale_wasm` binary (if it's not already downloaded).
3. You need to use the `wescale_wasm` binary to install the plugin.
    3.1 It will copy the wasm binary to the `WeScale` cluster. You can see the wasm binary in the `mysql.wasm_binary` system table.
    3.2 It will create a new filter and attach the wasm binary to it. You can see the filter using the `SHOW FILTERS` command.

You can specify detailed arguments for the `wescale_wasm` binary to install the plugin. 
```bash
# To see all the available arguments
./bin/wescale_wasm -h

# To specify the mysql arguments
./bin/wescale_wasm --command=install --mysql_host=127.0.0.1 --mysql_port=15306 --mysql_user=root --mysql_password=root

# To specify the wasm file arguments
./bin/wescale_wasm --command=install --wasm_file=./bin/my_plugin.wasm

# To specify the filter arguments
./bin/wescale_wasm --command=install --filter_name=foo --filter_desc='some kind of description' --filter_status=INACTIVE
```


## UnDeploy
**How to undeploy a wasm binary:**
1. You need to use the `wescale_wasm` binary to uninstall the plugin.
    1.1 It will delete the wasm binary from the `WeScale` cluster. You can see the wasm binary in the `mysql.wasm_binary` system table.
    1.2 It will delete the filter and detach the wasm binary from it. You can see the filter using the `SHOW FILTERS` command.
```bash
./bin/wescale_wasm --command=uninstall --filter_name=foo
```