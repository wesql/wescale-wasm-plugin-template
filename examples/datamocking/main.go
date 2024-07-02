package main

import (
	"github.com/earayu/sqlparser/go/sqltypes"
	"github.com/earayu/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
)

func main() {
	pkg.SetWasmPlugin(&DataMockingWasmPlugin{})
}

type DataMockingWasmPlugin struct {
}

func (a *DataMockingWasmPlugin) RunBeforeExecution() error {
	// do nothing
	return nil
}

func (a *DataMockingWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
	if queryResult == nil || errBefore != nil {
		return queryResult, errBefore
	}

	result := sqltypes.Proto3ToResult(queryResult)
	for _, row := range result.Rows {
		for _, value := range row {
			switch value.Type() {
			case sqltypes.VarChar, sqltypes.Char, sqltypes.Text:
				value = sqltypes.MakeTrusted(sqltypes.VarChar, []byte("foobar"))
			case sqltypes.Int8, sqltypes.Int16, sqltypes.Int24, sqltypes.Int32, sqltypes.Int64:
				value = sqltypes.NewInt8(42)
			case sqltypes.Uint8, sqltypes.Uint16, sqltypes.Uint24, sqltypes.Uint32, sqltypes.Uint64:
				value = sqltypes.NewUint64(42)
			}
		}
	}

	return queryResult, errBefore
}
