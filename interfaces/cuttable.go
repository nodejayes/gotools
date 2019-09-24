package interfaces

import "github.com/nodejayes/gotools/types"

type ICuttable interface {
	Clamp(lower, upper types.Number) *types.Number
}
