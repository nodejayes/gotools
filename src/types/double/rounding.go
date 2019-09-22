package double_type

import (
	"bytes"
	"fmt"
	"github.com/nodejayes/gotools/src/types/integer"
	"math"
)

// ceiling the Double by the given precision
func (d *Double) Ceil(precision int_type.Integer) Double {
	if precision.IsBelow(int_type.NewInteger(1)) {
		return NewDouble(math.Ceil(d.AsFloat64()))
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("1")
	for i := 0; i < precision.AsInt(); i++ {
		buf.WriteString("0")
	}
	factor := NewDouble(buf.String())
	converted := d.Multiply(factor)
	ceiling := NewDouble(math.Ceil(converted.AsFloat64()))
	return ceiling.Divide(factor)
}

// floor the Double by the given precision
func (d *Double) Floor(precision int_type.Integer) Double {
	if precision.IsBelow(int_type.NewInteger(1)) {
		return NewDouble(math.Floor(d.AsFloat64()))
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("1")
	for i := 0; i < precision.AsInt(); i++ {
		buf.WriteString("0")
	}
	factor := NewDouble(buf.String())
	converted := d.Multiply(factor)
	flooring := NewDouble(math.Floor(converted.AsFloat64()))
	return flooring.Divide(factor)
}

// round the Double by the given precision
func (d *Double) Round(precision int_type.Integer) Double {
	fmtBuf := bytes.NewBuffer([]byte{})
	fmtBuf.WriteString("%.")
	fmtBuf.WriteString(precision.ToString())
	fmtBuf.WriteString("f")
	return NewDouble(fmt.Sprintf(fmtBuf.String(), d.value))
}
