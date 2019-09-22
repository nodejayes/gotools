package double_type

import (
	"github.com/nodejayes/gotools/src/types/integer"
	"strconv"
)

// get the Double Value as float 32 bit
func (d *Double) AsFloat32() float32 {
	return float32(d.value)
}

// get the Double Value as float 64 bit
func (d *Double) AsFloat64() float64 {
	return d.value
}

// restricts the Double value to upper and lower
func (d *Double) Clamp(lower, upper Double) Double {
	if d.value < lower.value {
		return lower
	}
	if d.value > upper.value {
		return upper
	}
	return *d
}

// get a string represent the Double Value
func (d *Double) ToString(precision int_type.Integer) string {
	return strconv.FormatFloat(d.value, 'f', precision.AsInt(), 64)
}
