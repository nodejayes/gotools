package types

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// represent a Integer Data Type can be 16, 32, 64 bit
type Integer struct {
	value   int64
	isValid bool
}

// Creates a new Integer data type from the following types:
// int, int16, int32, int64
// float32, float64
// string, []byte
// Integer, Double
func NewInteger(v interface{}) Integer {
	switch v.(type) {
	case int:
		return Integer{
			value:   int64(v.(int)),
			isValid: true,
		}
	case int16:
		return Integer{
			value:   int64(v.(int16)),
			isValid: true,
		}
	case int32:
		return Integer{
			value:   int64(v.(int32)),
			isValid: true,
		}
	case int64:
		return Integer{
			value:   v.(int64),
			isValid: true,
		}
	case float32:
		return Integer{
			value:   int64(v.(float32)),
			isValid: true,
		}
	case float64:
		return Integer{
			value:   int64(v.(float64)),
			isValid: true,
		}
	case string:
		tmp, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			println("not supported value for Integer: " + v.(string))
			return Integer{
				value:   0,
				isValid: false,
			}
		}
		return Integer{
			value:   tmp,
			isValid: true,
		}
	case []byte:
		tmp, err := strconv.ParseInt(string(v.([]byte)), 10, 64)
		if err != nil {
			println("not supported value for Integer: " + string(v.([]byte)))
			return Integer{
				value:   0,
				isValid: false,
			}
		}
		return Integer{
			value:   tmp,
			isValid: true,
		}
	case Double:
		tmp := v.(Double)
		return Integer{
			value:   int64(tmp.AsFloat64()),
			isValid: true,
		}
	case Integer:
		return v.(Integer)
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for Integer")
	return Integer{
		value:   0,
		isValid: false,
	}
}

// Generates a new Integer random number in the range of min and max
func RandomInteger(min, max Integer) Integer {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewInteger(rand.Intn(max.AsInt()-min.AsInt()) + min.AsInt())
}

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
