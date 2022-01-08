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

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

//export proxy_on_request_headers
func proxyOnRequestHeaders(contextID uint32, numHeaders int, endOfStream bool) types.Action {
	ctx, ok := currentState.httpContexts[contextID]
	if !ok {
		panic("invalid context on proxy_on_request_headers")
	}

	currentState.setActiveContextID(contextID)
	return ctx.OnHttpRequestHeaders(numHeaders, endOfStream)
}

//export proxy_on_request_trailers
func proxyOnRequestTrailers(contextID uint32, numTrailers int) types.Action {
	ctx, ok := currentState.httpContexts[contextID]
	if !ok {
		panic("invalid context on proxy_on_request_trailers")
	}
	currentState.setActiveContextID(contextID)
	return ctx.OnHttpRequestTrailers(numTrailers)
}
