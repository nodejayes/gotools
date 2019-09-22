package types

import "strconv"

// get the Integer Value as int 64 bit
func (i *Integer) AsInt64() int64 {
	return i.value
}

// get the Integer Value as int 32 bit
func (i *Integer) AsInt32() int32 {
	return int32(i.value)
}

// get the Integer Value as int 16 bit
func (i *Integer) AsInt16() int16 {
	return int16(i.value)
}

// get the Integer Value as int
func (i *Integer) AsInt() int {
	return int(i.value)
}

// get a string represent the Integer Value
func (i *Integer) ToString() string {
	return strconv.FormatInt(i.value, 10)
}

// restricts the Integer value to upper and lower
func (i *Integer) Clamp(lower, upper Integer) Integer {
	if i.value < lower.value {
		return lower
	}
	if i.value > upper.value {
		return upper
	}
	return *i
}

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
func (d *Double) ToString(precision Integer) string {
	return strconv.FormatFloat(d.value, 'f', precision.AsInt(), 64)
}
