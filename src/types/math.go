package types

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

// addition of two Double values
func (d *Double) Add(v Double) Double {
	return NewDouble(d.value + v.value)
}

// subtraction of two Double values
func (d *Double) Subtract(v Double) Double {
	return NewDouble(d.value - v.value)
}

// multiplication of two Double values
func (d *Double) Multiply(v Double) Double {
	return NewDouble(d.value * v.value)
}

// division of two Double values
func (d *Double) Divide(v Double) Double {
	if v.value == 0 {
		panic("division by 0")
	}
	return NewDouble(d.value / v.value)
}
