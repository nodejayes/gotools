// The Types Package is Part of the GoTools Framework and defines/redefines some Basic DataTypes
package types

// represent a Integer Data Type can be 16, 32, 64 bit
type Integer struct {
	value   int64
	isValid bool
}

// represent a Double Data Type can be 32 or 64 bit
type Double struct {
	value   float64
	isValid bool
}
