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

package proxywasm

import (
	"fmt"
	"math"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/internal"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

// GetVMConfiguration is used for retrieving configurations given in the "vm_config.configuration" field.
// This hostcall is only available during types.PluginContext.OnVMStart call.
func GetVMConfiguration() ([]byte, error) {
	return getBuffer(internal.BufferTypeVMConfiguration, 0, math.MaxInt32)
}

// GetPluginConfiguration is used for retrieving configurations given in the "config.configuration" field.
// This hostcall is only available during types.PluginContext.OnPluginStart call.
func GetPluginConfiguration() ([]byte, error) {
	return getBuffer(internal.BufferTypePluginConfiguration, 0, math.MaxInt32)
}

func GetHttpRequestHeaders() ([][2]string, error) {
	return getMap(internal.MapTypeHttpRequestHeaders)
}

func ReplaceHttpRequestHeaders(headers [][2]string) error {
	return setMap(internal.MapTypeHttpRequestHeaders, headers)
}

func GetHttpRequestTrailers() ([][2]string, error) {
	return getMap(internal.MapTypeHttpRequestTrailers)
}

// LogTracef emit a message as a log with Trace log level.
func LogTrace(msg string) {
	internal.ProxyLog(internal.LogLevelTrace, internal.StringBytePtr(msg), len(msg))
}

// LogTracef formats according to a format specifier and emit as a log with Trace log level.
func LogTracef(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelTrace, internal.StringBytePtr(msg), len(msg))
}

// LogTracef emit a message as a log with Debug log level.
func LogDebug(msg string) {
	internal.ProxyLog(internal.LogLevelDebug, internal.StringBytePtr(msg), len(msg))
}

// LogDebugf formats according to a format specifier and emit as a log with Debug log level.
func LogDebugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelDebug, internal.StringBytePtr(msg), len(msg))
}

// LogTracef emit a message as a log with Info log level.
func LogInfo(msg string) {
	internal.ProxyLog(internal.LogLevelInfo, internal.StringBytePtr(msg), len(msg))
}

// LogInfof formats according to a format specifier and emit as a log with Info log level.
func LogInfof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelInfo, internal.StringBytePtr(msg), len(msg))
}

// LogTracef emit a message as a log with Warn log level.
func LogWarn(msg string) {
	internal.ProxyLog(internal.LogLevelWarn, internal.StringBytePtr(msg), len(msg))
}

// LogWarnf formats according to a format specifier and emit as a log with Warn log level.
func LogWarnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelWarn, internal.StringBytePtr(msg), len(msg))
}

// LogTracef emit a message as a log with Error log level.
func LogError(msg string) {
	internal.ProxyLog(internal.LogLevelError, internal.StringBytePtr(msg), len(msg))
}

// LogErrorf formats according to a format specifier and emit as a log with Error log level.
func LogErrorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelError, internal.StringBytePtr(msg), len(msg))
}

// LogTracef emit a message as a log with Critical log level.
func LogCritical(msg string) {
	internal.ProxyLog(internal.LogLevelCritical, internal.StringBytePtr(msg), len(msg))
}

// LogCriticalf formats according to a format specifier and emit as a log with Critical log level.
func LogCriticalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	internal.ProxyLog(internal.LogLevelCritical, internal.StringBytePtr(msg), len(msg))
}

func setMap(mapType internal.MapType, headers [][2]string) error {
	shs := internal.SerializeMap(headers)
	hp := &shs[0]
	hl := len(shs)
	return internal.StatusToError(internal.ProxySetHeaderMapPairs(mapType, hp, hl))
}

func getMapValue(mapType internal.MapType, key string) (string, error) {
	var rvs int
	var raw *byte
	if st := internal.ProxyGetHeaderMapValue(
		mapType, internal.StringBytePtr(key), len(key), &raw, &rvs,
	); st != internal.StatusOK {
		return "", internal.StatusToError(st)
	}

	ret := internal.RawBytePtrToString(raw, rvs)
	return ret, nil
}

func removeMapValue(mapType internal.MapType, key string) error {
	return internal.StatusToError(
		internal.ProxyRemoveHeaderMapValue(mapType, internal.StringBytePtr(key), len(key)),
	)
}

func replaceMapValue(mapType internal.MapType, key, value string) error {
	return internal.StatusToError(
		internal.ProxyReplaceHeaderMapValue(
			mapType, internal.StringBytePtr(key), len(key), internal.StringBytePtr(value), len(value),
		),
	)
}

func addMapValue(mapType internal.MapType, key, value string) error {
	return internal.StatusToError(
		internal.ProxyAddHeaderMapValue(
			mapType, internal.StringBytePtr(key), len(key), internal.StringBytePtr(value), len(value),
		),
	)
}

func getMap(mapType internal.MapType) ([][2]string, error) {
	var rvs int
	var raw *byte

	st := internal.ProxyGetHeaderMapPairs(mapType, &raw, &rvs)
	if st != internal.StatusOK {
		return nil, internal.StatusToError(st)
	} else if raw == nil {
		return nil, types.ErrorStatusNotFound
	}

	bs := internal.RawBytePtrToByteSlice(raw, rvs)
	return internal.DeserializeMap(bs), nil
}

func getBuffer(bufType internal.BufferType, start, maxSize int) ([]byte, error) {
	var retData *byte
	var retSize int
	switch st := internal.ProxyGetBufferBytes(bufType, start, maxSize, &retData, &retSize); st {
	case internal.StatusOK:
		if retData == nil {
			return nil, types.ErrorStatusNotFound
		}
		return internal.RawBytePtrToByteSlice(retData, retSize), nil
	default:
		return nil, internal.StatusToError(st)
	}
}

func appendToBuffer(bufType internal.BufferType, buffer []byte) error {
	var bufferData *byte
	if len(buffer) != 0 {
		bufferData = &buffer[0]
	}
	return internal.StatusToError(internal.ProxySetBufferBytes(bufType, math.MaxInt32, 0, bufferData, len(buffer)))
}

func prependToBuffer(bufType internal.BufferType, buffer []byte) error {
	var bufferData *byte
	if len(buffer) != 0 {
		bufferData = &buffer[0]
	}
	return internal.StatusToError(internal.ProxySetBufferBytes(bufType, 0, 0, bufferData, len(buffer)))
}

func replaceBuffer(bufType internal.BufferType, buffer []byte) error {
	var bufferData *byte
	if len(buffer) != 0 {
		bufferData = &buffer[0]
	}
	return internal.StatusToError(
		internal.ProxySetBufferBytes(bufType, 0, math.MaxInt32, bufferData, len(buffer)),
	)
}
