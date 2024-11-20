package main

import (
	"github.com/wesql/sqlparser/go/sqltypes"
	"github.com/wesql/sqlparser/go/vt/proto/query"
	"github.com/wesql/wescale-wasm-plugin-sdk/pkg"
	"math/rand"
	"time"
)

func main() {
	pkg.InitWasmPlugin(&DataMaskingWasmPlugin{})
}

type DataMaskingWasmPlugin struct {
}

func (a *DataMaskingWasmPlugin) RunBeforeExecution() error {
	// do nothing
	return nil
}

func (a *DataMaskingWasmPlugin) RunAfterExecution(queryResult *query.QueryResult, errBefore error) (*query.QueryResult, error) {
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, row := range result.Rows {
		mockResult.Rows[i] = make([]sqltypes.Value, len(mockResult.Fields))
		for j, field := range mockResult.Fields {
			switch field.Type {
			case query.Type_VARCHAR, query.Type_CHAR, query.Type_TEXT:
				mockResult.Rows[i][j] = sqltypes.MakeTrusted(field.Type, []byte("****"))
			case query.Type_INT8, query.Type_INT16, query.Type_INT24, query.Type_INT32, query.Type_INT64:
				mockResult.Rows[i][j] = sqltypes.NewInt8(int8(r.Intn(200)))
			case query.Type_UINT8, query.Type_UINT16, query.Type_UINT24, query.Type_UINT32, query.Type_UINT64:
				mockResult.Rows[i][j] = sqltypes.NewInt8(int8(r.Intn(200)))
			default:
				mockResult.Rows[i][j] = row[j]
			}
		}
	}

	return sqltypes.ResultToProto3(mockResult), errBefore
}
