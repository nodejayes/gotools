package int_type

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
