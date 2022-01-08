package sdk

import (
	"strconv"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type SdkVmContext struct {
	types.DefaultVMContext

	newPluginContext func() types.PluginContext
}

func NewColumnMappingPlugin(handler ColumnMappingHandler) *SdkVmContext {
	f := func() types.PluginContext {
		return newColumnMappingPluginContext(handler)
	}
	return newSdkVmContext(f)
}

func newSdkVmContext(pluginCtxFactory func() types.PluginContext) *SdkVmContext {
	return &SdkVmContext{
		newPluginContext: pluginCtxFactory,
	}
}

func (s *SdkVmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return s.newPluginContext()
}

func toValueMap(hs [][2]string) map[int]string {
	ret := make(map[int]string)
	for _, h := range hs {
		colIdx, _ := strconv.Atoi(h[0])
		ret[colIdx] = h[1]
	}
	return ret
}
