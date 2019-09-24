package interfaces

import "github.com/nodejayes/gotools/types"

type IComparable interface {
	Equals(v types.Number) bool
}
