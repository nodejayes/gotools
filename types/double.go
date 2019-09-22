// The Types Package is Part of the GoTools Framework and defines/redefines some Basic DataTypes
package types

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// represent a Double Data Type can be 32 or 64 bit
type Double struct {
	value   float64
	isValid bool
}

// Creates a new Double data type from the following types:
// int, int16, int32, int64
// float32, float64
// string, []byte
// Integer, Double
func NewDouble(v interface{}) Double {
	switch v.(type) {
	case int:
		return Double{
			value:   float64(v.(int)),
			isValid: true,
		}
	case int16:
		return Double{
			value:   float64(v.(int16)),
			isValid: true,
		}
	case int32:
		return Double{
			value:   float64(v.(int32)),
			isValid: true,
		}
	case int64:
		return Double{
			value:   float64(v.(int64)),
			isValid: true,
		}
	case float32:
		return Double{
			value:   float64(v.(float32)),
			isValid: true,
		}
	case float64:
		return Double{
			value:   v.(float64),
			isValid: true,
		}
	case string:
		tmp, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			println("invalid value for Double: " + v.(string))
			return Double{
				value:   0,
				isValid: false,
			}
		}
		return Double{
			value:   tmp,
			isValid: true,
		}
	case []byte:
		tmp, err := strconv.ParseFloat(string(v.([]byte)), 64)
		if err != nil {
			println("invalid value for Double: " + string(v.([]byte)))
			return Double{
				value:   0,
				isValid: false,
			}
		}
		return Double{
			value:   tmp,
			isValid: true,
		}
	case Integer:
		tmp := v.(Integer)
		return Double{
			value:   float64(tmp.AsInt()),
			isValid: true,
		}
	case Double:
		return v.(Double)
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for Double")
	return Double{
		value:   0,
		isValid: false,
	}
}

// Generates a new Double random number in the range of min and max
func RandomDouble(min, max Double) Double {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewDouble(min.AsFloat64() + rand.Float64()*(max.AsFloat64()-min.AsFloat64()))
}

// ceiling the Double by the given precision
func (d *Double) Ceil(precision Integer) Double {
	if precision.IsBelow(NewInteger(1)) {
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
func (d *Double) Floor(precision Integer) Double {
	if precision.IsBelow(NewInteger(1)) {
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
func (d *Double) Round(precision Integer) Double {
	fmtBuf := bytes.NewBuffer([]byte{})
	fmtBuf.WriteString("%.")
	fmtBuf.WriteString(precision.ToString())
	fmtBuf.WriteString("f")
	return NewDouble(fmt.Sprintf(fmtBuf.String(), d.value))
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
