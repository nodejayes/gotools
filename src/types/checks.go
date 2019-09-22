package types

// is the Value a Valid Integer
func (i *Integer) IsValid() bool {
	return i.isValid
}

// check if the Integer is between given start and end including the borders start and end
func (i *Integer) IsInRange(start, end Integer) bool {
	return i.value <= end.value && i.value >= start.value
}

// check if the given Integer the same Value as this Integer
func (i *Integer) Equals(v Integer) bool {
	return i.value == v.value
}

// check if the given Integer is bigger as this Integer not including the Border
func (i *Integer) IsAbove(v Integer) bool {
	return i.value > v.value
}

// check if the given Integer is smaller as this Integer not including the Border
func (i *Integer) IsBelow(v Integer) bool {
	return i.value < v.value
}

// is the Value a Valid Double
func (d *Double) IsValid() bool {
	return d.isValid
}

// check if the Double is between given start and end including the borders start and end
func (d *Double) IsInRange(start, end Double) bool {
	return d.value <= end.value && d.value >= start.value
}

// check if the given Integer the same Value as this Double
func (d *Double) Equals(v Double) bool {
	return d.value == v.value
}

// check if the given Double is bigger as this Integer not including the Border
func (d *Double) IsAbove(v Double) bool {
	return d.value > v.value
}

// check if the given Double is smaller as this Integer not including the Border
func (d *Double) IsBelow(v Double) bool {
	return d.value < v.value
}
