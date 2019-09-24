package interfaces

import "github.com/nodejayes/gotools/types"

type IRangeCheck interface {
	IsInRange(start, end types.Number) bool
	IsAbove(v types.Number) bool
	IsBelow(v types.Number) bool
}
