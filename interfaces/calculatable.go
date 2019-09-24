package interfaces

import "github.com/nodejayes/gotools/types"

type ICalculatable interface {
	Add(v types.Number) *types.Number
	Subtract(v types.Number) *types.Number
	Multiply(v types.Number) *types.Number
	Divide(v types.Number) *types.Number
}
