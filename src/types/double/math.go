package double_type

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
