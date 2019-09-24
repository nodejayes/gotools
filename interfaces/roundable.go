package interfaces

import "github.com/nodejayes/gotools/types"

type IRoundable interface {
	Round(precision types.Number) *types.Number
}

type IFloorable interface {
	Floor(precision types.Number) *types.Number
}

type ICeilable interface {
	Ceil(precision types.Number) *types.Number
}
