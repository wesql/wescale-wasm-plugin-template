# WeScale-Wasm-Plugin-Template

## Pre-requisites
* You need to have TinyGo to compile the code, you can install it by following the instructions [here](https://tinygo.org/getting-started/install/).

* You will also need to have the `wescale_wasm` binary to deploy the plugin, you can download it by running the following command.
```bash
make install-wescale-wasm
```

## Write your code
You can write your code in the `main.go` file. You can use the following template to write your code.
You can see the example code in the `examples` directory.

## Build
You can build the wasm binary by running the following command. You can specify the wasm file name by using the WASM_FILE variable. 
You can see the wasm binary in the bin directory.
```bash
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
    * It will copy the wasm binary to the `WeScale` cluster. You can see the wasm binary in the `mysql.wasm_binary` system table.
    * It will create a new filter and attach the wasm binary to it. You can see the filter using the `SHOW FILTERS` command.

You can specify detailed arguments for the `wescale_wasm` binary to install the plugin. 
```bash
# To see all the available arguments
./bin/wescale_wasm -h

# To specify the mysql arguments
./bin/wescale_wasm --command=install --wasm_file=./bin/my_plugin.wasm --mysql_host=127.0.0.1 --mysql_port=15306 --mysql_user=root --mysql_password=root

# To specify the filter arguments
./bin/wescale_wasm --command=install --wasm_file=./bin/my_plugin.wasm --filter_name=my_plugin_wasm_filter --filter_desc='some kind of description' --filter_status=INACTIVE
```


## UnDeploy
**How to undeploy a wasm binary:**
1. You need to use the `wescale_wasm` binary to uninstall the plugin.
    * It will delete the wasm binary from the `WeScale` cluster. You can see the wasm binary in the `mysql.wasm_binary` system table.
    * It will delete the filter and detach the wasm binary from it. You can see the filter using the `SHOW FILTERS` command.
```bash
./bin/wescale_wasm --command=uninstall --filter_name=my_plugin_wasm_filter
```