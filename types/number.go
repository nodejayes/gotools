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

// represent a Numeric Value int, int16, int32, int64, float32 or float64
type Number struct {
	value   float64
	isValid bool
}

func NewNumber(v interface{}) *Number {
	switch v.(type) {
	case int:
		return &Number{
			value:   float64(v.(int)),
			isValid: true,
		}
	case int16:
		return &Number{
			value:   float64(v.(int16)),
			isValid: true,
		}
	case int32:
		return &Number{
			value:   float64(v.(int32)),
			isValid: true,
		}
	case int64:
		return &Number{
			value:   float64(v.(int64)),
			isValid: true,
		}
	case float32:
		return &Number{
			value:   float64(v.(float32)),
			isValid: true,
		}
	case float64:
		return &Number{
			value:   v.(float64),
			isValid: true,
		}
	case string:
		tmp := stringToNumber(v.(string))
		return &tmp
	case []byte:
		tmp := stringToNumber(string(v.([]byte)))
		return &tmp
	case Number:
		tmp := v.(Number)
		return &tmp
	case *Number:
		tmp := v.(*Number)
		return tmp
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for Double")
	return &Number{
		value:   0,
		isValid: false,
	}
}

// Generates a new random number in the range of min and max in int precision
func RandomNumberInt(min, max Number) *Number {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewNumber(min.AsInt() + rand.Intn(max.AsInt()-min.AsInt()))
}

// Generates a new random number in the range of min and max in float precision
func RandomNumberFloat(min, max Number) *Number {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewNumber(min.AsFloat64() + rand.Float64()*(max.AsFloat64()-min.AsFloat64()))
}

func stringToNumber(v string) Number {
	tmp, err := strconv.ParseFloat(v, 64)
	if err != nil {
		println("invalid value for Double: " + v)
		return Number{
			value:   0,
			isValid: false,
		}
	}
	return Number{
		value:   tmp,
		isValid: true,
	}
}

// get the Number Value as int 64 bit
func (i *Number) AsInt64() int64 {
	return int64(i.value)
}

// get the Number Value as int 32 bit
func (i *Number) AsInt32() int32 {
	return int32(i.value)
}

// get the Number Value as int 16 bit
func (i *Number) AsInt16() int16 {
	return int16(i.value)
}

// get the Number Value as int
func (i *Number) AsInt() int {
	return int(i.value)
}

// get the Double Value as float 32 bit
func (i *Number) AsFloat32() float32 {
	return float32(i.value)
}

// get the Double Value as float 64 bit
func (i *Number) AsFloat64() float64 {
	return i.value
}

// get a string represent the Number Value
func (i *Number) AsString(precision Number) string {
	return strconv.FormatFloat(i.value, 'f', precision.AsInt(), 64)
}

// ge the string representation as Byte Array
func (i *Number) AsByte(precision Number) []byte {
	return []byte(i.AsString(precision))
}

// restricts the Number value to upper and lower
func (i *Number) Clamp(lower, upper Number) *Number {
	if i.value < lower.value {
		return &lower
	}
	if i.value > upper.value {
		return &upper
	}
	return i
}

// is the Value a Valid Number
func (i *Number) IsValid() bool {
	return i.isValid
}

// check if the Number is between given start and end including the borders start and end
func (i *Number) IsInRange(start, end Number) bool {
	return i.value <= end.value && i.value >= start.value
}

// check if the given Number the same Value as this Number
func (i *Number) Equals(v Number) bool {
	return i.value == v.value
}

// check if the given Number is bigger as this Number not including the Border
func (i *Number) IsAbove(v Number) bool {
	return i.value > v.value
}

// check if the given Number is smaller as this Number not including the Border
func (i *Number) IsBelow(v Number) bool {
	return i.value < v.value
}

// addition of two Number values
func (i *Number) Add(v Number) *Number {
	return NewNumber(i.value + v.value)
}

// subtraction of two Number values
func (i *Number) Subtract(v Number) *Number {
	return NewNumber(i.value - v.value)
}

// multiplication of two Number values
func (i *Number) Multiply(v Number) *Number {
	return NewNumber(i.value * v.value)
}

// division of two Number values
func (i *Number) Divide(v Number) *Number {
	if v.value == 0 {
		return NewNumber(i.value / 1)
	}
	return NewNumber(i.value / v.value)
}

// ceiling the Number by the given precision
func (i *Number) Ceil(precision Number) *Number {
	if precision.IsBelow(*NewNumber(1)) {
		return NewNumber(math.Ceil(i.AsFloat64()))
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("1")
	for i := 0; i < precision.AsInt(); i++ {
		buf.WriteString("0")
	}
	factor := NewNumber(buf.String())
	converted := i.Multiply(*factor)
	ceiling := NewNumber(math.Ceil(converted.AsFloat64()))
	return ceiling.Divide(*factor)
}

// floor the Number by the given precision
func (i *Number) Floor(precision Number) *Number {
	if precision.IsBelow(*NewNumber(1)) {
		return NewNumber(math.Floor(i.AsFloat64()))
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("1")
	for i := 0; i < precision.AsInt(); i++ {
		buf.WriteString("0")
	}
	factor := NewNumber(buf.String())
	converted := i.Multiply(*factor)
	flooring := NewNumber(math.Floor(converted.AsFloat64()))
	return flooring.Divide(*factor)
}

// round the Number by the given precision
func (i *Number) Round(precision Number) *Number {
	fmtBuf := bytes.NewBuffer([]byte{})
	fmtBuf.WriteString("%.")
	fmtBuf.WriteString(precision.AsString(*NewNumber(0)))
	fmtBuf.WriteString("f")
	return NewNumber(fmt.Sprintf(fmtBuf.String(), i.value))
}
