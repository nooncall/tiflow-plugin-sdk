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

//go:build !proxytest

package internal

//export proxy_log
func ProxyLog(logLevel LogLevel, messageData *byte, messageSize int) Status

//export proxy_get_header_map_value
func ProxyGetHeaderMapValue(mapType MapType, keyData *byte, keySize int, returnValueData **byte, returnValueSize *int) Status

//export proxy_add_header_map_value
func ProxyAddHeaderMapValue(mapType MapType, keyData *byte, keySize int, valueData *byte, valueSize int) Status

//export proxy_replace_header_map_value
func ProxyReplaceHeaderMapValue(mapType MapType, keyData *byte, keySize int, valueData *byte, valueSize int) Status

//export proxy_remove_header_map_value
func ProxyRemoveHeaderMapValue(mapType MapType, keyData *byte, keySize int) Status

//export proxy_get_header_map_pairs
func ProxyGetHeaderMapPairs(mapType MapType, returnValueData **byte, returnValueSize *int) Status

//export proxy_set_header_map_pairs
func ProxySetHeaderMapPairs(mapType MapType, mapData *byte, mapSize int) Status

//export proxy_get_buffer_bytes
func ProxyGetBufferBytes(bufferType BufferType, start int, maxSize int, returnBufferData **byte, returnBufferSize *int) Status

//export proxy_set_buffer_bytes
func ProxySetBufferBytes(bufferType BufferType, start int, maxSize int, bufferData *byte, bufferSize int) Status

//export proxy_set_effective_context
func ProxySetEffectiveContext(contextID uint32) Status
