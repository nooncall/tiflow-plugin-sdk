package internal

import (
	"errors"
	"strconv"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type BufferType uint32

const (
	BufferTypeVMConfiguration     BufferType = 6
	BufferTypePluginConfiguration BufferType = 7
)

type LogLevel uint32

const (
	LogLevelTrace    LogLevel = 0
	LogLevelDebug    LogLevel = 1
	LogLevelInfo     LogLevel = 2
	LogLevelWarn     LogLevel = 3
	LogLevelError    LogLevel = 4
	LogLevelCritical LogLevel = 5
	LogLevelMax      LogLevel = 6
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelTrace:
		return "trace"
	case LogLevelDebug:
		return "debug"
	case LogLevelInfo:
		return "info"
	case LogLevelWarn:
		return "warn"
	case LogLevelError:
		return "error"
	case LogLevelCritical:
		return "critical"
	default:
		panic("invalid log level")
	}
}

type MapType uint32

const (
	MapTypeHttpRequestHeaders  MapType = 0
	MapTypeHttpRequestTrailers MapType = 1
)

type MetricType uint32

const (
	MetricTypeCounter   = 0
	MetricTypeGauge     = 1
	MetricTypeHistogram = 2
)

type StreamType uint32

const (
	StreamTypeRequest    StreamType = 0
	StreamTypeResponse   StreamType = 1
	StreamTypeDownstream StreamType = 2
	StreamTypeUpstream   StreamType = 3
)

type Status uint32

const (
	StatusOK              Status = 0
	StatusNotFound        Status = 1
	StatusBadArgument     Status = 2
	StatusEmpty           Status = 7
	StatusCasMismatch     Status = 8
	StatusInternalFailure Status = 10
	StatusUnimplemented   Status = 12
)

//go:inline
func StatusToError(status Status) error {
	switch Status(status) {
	case StatusOK:
		return nil
	case StatusNotFound:
		return types.ErrorStatusNotFound
	case StatusBadArgument:
		return types.ErrorStatusBadArgument
	case StatusEmpty:
		return types.ErrorStatusEmpty
	case StatusCasMismatch:
		return types.ErrorStatusCasMismatch
	case StatusInternalFailure:
		return types.ErrorInternalFailure
	case StatusUnimplemented:
		return types.ErrorUnimplemented
	}
	return errors.New("unknown status code: " + strconv.Itoa(int(status)))
}
