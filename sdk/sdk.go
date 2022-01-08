package sdk

import (
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
