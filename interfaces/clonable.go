package interfaces

import "github.com/nodejayes/gotools/types"

type IClonableNumber interface {
	Clone() *types.Number
}
