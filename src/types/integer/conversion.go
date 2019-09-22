package int_type

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
