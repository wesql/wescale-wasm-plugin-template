package main

import (
	"github.com/wesql/sqlparser/go/sqltypes"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg/types"
)

func main() {
	pkg.InitWasmPlugin(&DataMockingWasmPlugin{})
}

type DataMockingWasmPlugin struct {
}

func (a *DataMockingWasmPlugin) RunBeforeExecution(pluginCtx types.WasmPluginContext) error {
	// do nothing
	return nil
}

func (a *DataMockingWasmPlugin) RunAfterExecution(pluginCtx types.WasmPluginContext, queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult == nil || errBefore != nil {
		return queryResult, errBefore
	}

	result := sqltypes.Proto3ToResult(queryResult)
	if result.Fields == nil || len(result.Fields) == 0 {
		return queryResult, errBefore
	}
	mockResult := sqltypes.MakeTestResult(result.Fields)
	if len(result.Rows) > 0 {
		mockResult.Rows = make([][]sqltypes.Value, len(result.Rows))
	}
	for i, row := range result.Rows {
		mockResult.Rows[i] = make([]sqltypes.Value, len(mockResult.Fields))
		for j, field := range mockResult.Fields {
			switch field.Type {
			case query.Type_VARCHAR, query.Type_CHAR, query.Type_TEXT:
				mockResult.Rows[i][j] = sqltypes.MakeTrusted(field.Type, []byte("mocked"))
			case query.Type_INT8, query.Type_INT16, query.Type_INT24, query.Type_INT32, query.Type_INT64:
				mockResult.Rows[i][j] = sqltypes.MakeTrusted(field.Type, []byte("42"))
			case query.Type_UINT8, query.Type_UINT16, query.Type_UINT24, query.Type_UINT32, query.Type_UINT64:
				mockResult.Rows[i][j] = sqltypes.MakeTrusted(field.Type, []byte("42"))
			default:
				mockResult.Rows[i][j] = row[j]
			}
		}
	}

	return sqltypes.ResultToProto3(mockResult), errBefore
}
