package main

import (
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func main() {
	pkg.InitWasmPlugin(&CustomWasmPlugin{})
}

// TODO: RENAME THIS STRUCT TO YOUR PLUGIN NAME
type CustomWasmPlugin struct {
}

func (a *CustomWasmPlugin) RunBeforeExecution(pluginCtx types.WasmPluginContext) error {
	//TODO: Implement your logic here
	return nil
}

func (a *CustomWasmPlugin) RunAfterExecution(pluginCtx types.WasmPluginContext, queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	//TODO: Implement your logic here
	return queryResult, errBefore
}
