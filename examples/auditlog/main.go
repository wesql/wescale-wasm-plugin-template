package main

import (
	"fmt"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	hostfunction "github.com/wesql/wescale-wasm-plugin-sdk/pkg/host_functions"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func main() {
	pkg.InitWasmPlugin(&AuditLogWasmPlugin{})
}

type AuditLogWasmPlugin struct {
}

func (a *AuditLogWasmPlugin) RunBeforeExecution(pluginCtx types.WasmPluginContext) error {
	query, err := hostfunction.GetHostQuery(pluginCtx)
	if err != nil {
		return err
	}
	hostfunction.InfoLog(pluginCtx, "execute SQL: "+query)
	return nil
}

func (a *AuditLogWasmPlugin) RunAfterExecution(pluginCtx types.WasmPluginContext, queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult != nil {
		hostfunction.InfoLog(pluginCtx, fmt.Sprintf("returned rows: %v", len(queryResult.Rows)))
		hostfunction.InfoLog(pluginCtx, fmt.Sprintf("affected rows: %v", queryResult.RowsAffected))
	}
	if errBefore != nil {
		hostfunction.InfoLog(pluginCtx, "execution error: "+errBefore.Error())
	}

	return queryResult, errBefore
}
