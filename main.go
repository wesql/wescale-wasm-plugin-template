package main

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
)

func main() {
	pkg.SetWasmPlugin(&CustomWasmPlugin{})
}

// TODO: RENAME THIS STRUCT TO YOUR PLUGIN NAME
type CustomWasmPlugin struct {
}

func (a *CustomWasmPlugin) RunBeforeExecution() error {
	//TODO: Implement your logic here
	return nil
}

func (a *CustomWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	//TODO: Implement your logic here
	return queryResult, errBefore
}
