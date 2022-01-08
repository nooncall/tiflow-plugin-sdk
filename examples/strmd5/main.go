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

package main

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/sdk"
)

func main() {
	handler := func(s string) string {
		hasher := md5.New()
		hasher.Write([]byte(s))
		newVal := hex.EncodeToString(hasher.Sum(nil))
		return newVal
	}
	vmContext := sdk.NewColumnMappingPlugin(handler)
	proxywasm.SetVMContext(vmContext)
}
