package int_type

// addition of two Integer values
func (i *Integer) Add(v Integer) Integer {
	return NewInteger(i.value + v.value)
}

// subtraction of two Integer values
func (i *Integer) Subtract(v Integer) Integer {
	return NewInteger(i.value - v.value)
}

// multiplication of two Integer values
func (i *Integer) Multiply(v Integer) Integer {
	return NewInteger(i.value * v.value)
}

// division of two Integer values
func (i *Integer) Divide(v Integer) Integer {
	if v.value == 0 {
		panic("division by 0")
	}
	return NewInteger(i.value / v.value)
}
