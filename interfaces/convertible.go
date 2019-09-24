package interfaces

import "github.com/nodejayes/gotools/types"

type IIntConvertible interface {
	AsInt() int
	AsInt16() int16
	AsInt32() int32
	AsInt64() int64
}

type IFloatConvertible interface {
	AsFloat32() float32
	AsFloat64() float64
}

type IStringConvertible interface {
	AsString(precision types.Number) string
}

type IByteConvertible interface {
	AsByte(precision types.Number) []byte
}
