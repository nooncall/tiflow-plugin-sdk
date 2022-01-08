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

package internal

import "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"

//nolint
//export proxy_on_memory_allocate
func proxyOnMemoryAllocate(size uint) *byte {
	buf := make([]byte, size)
	return &buf[0]
}

//export proxy_on_vm_start
func proxyOnVMStart(_ uint32, vmConfigurationSize int) types.OnVMStartStatus {
	return currentState.vmContext.OnVMStart(vmConfigurationSize)
}

//export proxy_on_configure
func proxyOnConfigure(pluginContextID uint32, pluginConfigurationSize int) types.OnPluginStartStatus {
	ctx, ok := currentState.pluginContexts[pluginContextID]
	if !ok {
		panic("invalid context on proxy_on_configure")
	}
	currentState.setActiveContextID(pluginContextID)
	return ctx.context.OnPluginStart(pluginConfigurationSize)
}

//export proxy_on_context_create
func proxyOnContextCreate(contextID uint32, pluginContextID uint32) {
	if pluginContextID == 0 {
		currentState.createPluginContext(contextID)
	} else if currentState.createHttpContext(contextID, pluginContextID) {
	} else if currentState.createTcpContext(contextID, pluginContextID) {
	} else {
		panic("invalid context id on proxy_on_context_create")
	}
}

//export proxy_on_log
func proxyOnLog(contextID uint32) {
	if ctx, ok := currentState.tcpContexts[contextID]; ok {
		currentState.setActiveContextID(contextID)
		ctx.OnStreamDone()
	} else if ctx, ok := currentState.httpContexts[contextID]; ok {
		currentState.setActiveContextID(contextID)
		ctx.OnHttpStreamDone()
	}
}

//export proxy_on_done
func proxyOnDone(contextID uint32) bool {
	if ctx, ok := currentState.pluginContexts[contextID]; ok {
		currentState.setActiveContextID(contextID)
		return ctx.context.OnPluginDone()
	}
	return true
}

//export proxy_on_delete
func proxyOnDelete(contextID uint32) {
	delete(currentState.contextIDToRootID, contextID)
	if _, ok := currentState.tcpContexts[contextID]; ok {
		delete(currentState.tcpContexts, contextID)
	} else if _, ok = currentState.httpContexts[contextID]; ok {
		delete(currentState.httpContexts, contextID)
	} else if _, ok = currentState.pluginContexts[contextID]; ok {
		delete(currentState.pluginContexts, contextID)
	} else {
		panic("invalid context on proxy_on_delete")
	}
}

//export proxy_on_tick
func proxyOnTick(pluginContextID uint32) {
	ctx, ok := currentState.pluginContexts[pluginContextID]
	if !ok {
		panic("invalid root_context_id")
	}
	currentState.setActiveContextID(pluginContextID)
	ctx.context.OnTick()
}

//nolint
//export proxy_abi_version_0_2_0
func proxyABIVersion() {}
