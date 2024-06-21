package main

import (
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	hostfunction "github.com/wesql/wescale-wasm-plugin-sdk/pkg/host_functions/v1alpha1"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/proto/query"
)

func main() {
	pkg.SetWasmPlugin(&CustomWasmPlugin{})
}

// TODO: RENAME THIS STRUCT TO YOUR PLUGIN NAME
type CustomWasmPlugin struct {
}

func (a *CustomWasmPlugin) RunBeforeExecution() error {
	hostfunction.GetQueryResult()
	//TODO: Implement your logic here
	return nil
}

func (a *CustomWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	//TODO: Implement your logic here
	return queryResult, errBefore
}
