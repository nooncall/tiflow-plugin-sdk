// Copyright 2020-2021 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

// VMContext corresponds to each Wasm VM machine and its configuration. Thefore,
// this is the entrypoint for extending your network proxy.
// Its lifetime is exactly the same as Wasm Virtual Machines on the host.
type VMContext interface {
	// OnVMStart is called after the VM is created and main function is called.
	// During this call, GetVMConfiguration hostcall is available and can be used to
	// retrieve the configuration set at vm_config.configuration.
	// This is mainly used for doing Wasm VM-wise initialization.
	OnVMStart(vmConfigurationSize int) OnVMStartStatus

	// NewPluginContext is used for creating PluginContext for each plugin configurations.
	NewPluginContext(contextID uint32) PluginContext
}

// PluginContext corresponds to each different plugin configurations (config.configuration).
// Each configuration is usually given at each http/tcp filter in a listener in the hosts,
// so PluginContext is responsible for creating "filter instances" for each Tcp/Http streams on the listener.
type PluginContext interface {
	// OnPluginStart is called on all plugin contexts (after OnVmStart if this is the VM context).
	// During this call, GetPluginConfiguration is available and can be used to
	// retrieve the configuration set at config.configuration in envoy.yaml
	OnPluginStart(pluginConfigurationSize int) OnPluginStartStatus

	// onPluginDone is called right before plugin contexts are deleted by hosts.
	// Return false to indicate it's in a pending state to do some more work left.
	// In that case, must call PluginDone() host call after the work is done to indicate that
	// hosts can kill this contexts.
	OnPluginDone() bool

	// OnTick is called when SetTickPeriodMilliSeconds hostcall is called by this plugin context.
	// This can be used for doing some asynchronous tasks in parallel to stream processing.
	OnTick()

	// NewHttpContext is used for creating HttpContext for each Http streams.
	// Return nil to indicate this PluginContext is not for HttpContext.
	NewHttpContext(contextID uint32) HttpContext
}

// HttpContext corresponds to each Http stream and is created by PluginContext via NewHttpContext.
type HttpContext interface {
	// OnHttpRequestHeaders is called when request headers arrives.
	// Return types.ActionPause if you want to stop sending headers to upstream.
	OnHttpRequestHeaders(numHeaders int, endOfStream bool) Action

	// OnHttpRequestTrailers is called when request trailers arrives.
	// Return types.ActionPause if you want to stop sending trailers to upstream.
	OnHttpRequestTrailers(numTrailers int) Action

	// OnHttpStreamDone is called before the host deletes this context.
	// You can retreive the HTTP request/response information (such headers, etc.) during this calls.
	// This can be used for implementing logging feature.
	OnHttpStreamDone()
}

// DefaultContexts are no-op implementation of contexts.
// Users can embed them into their custom contexts, so that
// they only have to implement methods they want.
type (
	// DefaultVMContext provides the no-op implementation of VMContext interface.
	DefaultVMContext struct{}

	// DefaultPluginContext provides the no-op implementation of PluginContext interface.
	DefaultPluginContext struct{}

	// DefaultTcpContext provides the no-op implementation of TcpContext interface.
	DefaultTcpContext struct{}

	// DefaultHttpContext provides the no-op implementation of HttpContext interface.
	DefaultHttpContext struct{}
)

// impl VMContext
func (*DefaultVMContext) OnVMStart(vmConfigurationSize int) OnVMStartStatus { return OnVMStartStatusOK }
func (*DefaultVMContext) NewPluginContext(contextID uint32) PluginContext {
	return &DefaultPluginContext{}
}

// impl PluginContext
func (*DefaultPluginContext) OnQueueReady(uint32) {}
func (*DefaultPluginContext) OnTick()             {}
func (*DefaultPluginContext) OnPluginStart(int) OnPluginStartStatus {
	return OnPluginStartStatusOK
}
func (*DefaultPluginContext) OnPluginDone() bool                { return true }
func (*DefaultPluginContext) NewHttpContext(uint32) HttpContext { return nil }

// impl HttpContext
func (*DefaultHttpContext) OnHttpRequestHeaders(int, bool) Action  { return ActionContinue }
func (*DefaultHttpContext) OnHttpRequestBody(int, bool) Action     { return ActionContinue }
func (*DefaultHttpContext) OnHttpRequestTrailers(int) Action       { return ActionContinue }
func (*DefaultHttpContext) OnHttpResponseHeaders(int, bool) Action { return ActionContinue }
func (*DefaultHttpContext) OnHttpResponseBody(int, bool) Action    { return ActionContinue }
func (*DefaultHttpContext) OnHttpResponseTrailers(int) Action      { return ActionContinue }
func (*DefaultHttpContext) OnHttpStreamDone()                      {}

var (
	_ VMContext     = &DefaultVMContext{}
	_ PluginContext = &DefaultPluginContext{}
	_ HttpContext   = &DefaultHttpContext{}
)
