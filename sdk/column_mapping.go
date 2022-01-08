package sdk

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type ColumnMappingHandler func(val string) string

type columnMappingPluginContext struct {
	types.DefaultPluginContext

	handle ColumnMappingHandler
}

type columnMappingReqContext struct {
	types.DefaultHttpContext
	contextID uint32
	handle    ColumnMappingHandler
}

func newColumnMappingPluginContext(handler ColumnMappingHandler) types.PluginContext {
	return &columnMappingPluginContext{
		handle: handler,
	}
}

func newColumnMappingReqContext(handler ColumnMappingHandler) types.HttpContext {
	return &columnMappingReqContext{
		handle: handler,
	}
}

func (c *columnMappingPluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return newColumnMappingReqContext(c.handle)
}

func (c *columnMappingReqContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get values: %v", err)
	}
	//for _, h := range hs {
	//	proxywasm.LogInfof("before set: request header --> %s: %s", h[0], h[1])
	//}

	colIndexes, err := proxywasm.GetHttpRequestTrailers()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request trailers: %v", err)
	}
	//proxywasm.LogInfof("len request trailers: %d", len(colIndexes))
	for i, colIdx := range colIndexes {
		proxywasm.LogInfof("colIdx --> i: %d, idx: %s", i, colIdx[0])
	}

	for j := 0; j < len(hs); j++ {
		for i := 0; i < len(colIndexes); i++ {
			if hs[j][0] == colIndexes[i][0] {
				hs[j][1] = c.handle(hs[j][1])
			}
		}
	}

	if err := proxywasm.ReplaceHttpRequestHeaders(hs); err != nil {
		proxywasm.LogCriticalf("failed to set request headers, new: %v, err: %v", hs, err)
	}

	for _, h := range hs {
		proxywasm.LogInfof("after set: request header --> %s: %s", h[0], h[1])
	}
	return types.ActionContinue
}
