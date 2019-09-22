package double_type

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
