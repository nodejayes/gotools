package types

import (
	"reflect"
	"strconv"
)

type DivisibleNumber struct {
	value   float64
	isValid bool
}

func NewDivisibleNumber(v interface{}) DivisibleNumber {
	switch v.(type) {
	case int:
		tmp := v.(int)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp),
			isValid: true,
		}
	case int16:
		tmp := v.(int16)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp),
			isValid: true,
		}
	case int32:
		tmp := v.(int32)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp),
			isValid: true,
		}
	case int64:
		tmp := v.(int64)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp),
			isValid: true,
		}
	case float32:
		tmp := v.(float32)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp),
			isValid: true,
		}
	case float64:
		tmp := v.(float64)
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   tmp,
			isValid: true,
		}
	case string:
		tmp, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			println("invalid value for Double: " + v.(string))
			return DivisibleNumber{
				value:   0,
				isValid: false,
			}
		}
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   tmp,
			isValid: true,
		}
	case []byte:
		tmp, err := strconv.ParseFloat(string(v.([]byte)), 64)
		if err != nil {
			println("invalid value for Double: " + string(v.([]byte)))
			return DivisibleNumber{
				value:   0,
				isValid: false,
			}
		}
		if tmp == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   tmp,
			isValid: true,
		}
	case Integer:
		tmp := v.(Integer)
		if tmp.AsInt() == 0 {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   float64(tmp.AsInt()),
			isValid: true,
		}
	case Double:
		tmp := v.(Double)
		if tmp.AsFloat64() == float64(0) {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return DivisibleNumber{
			value:   tmp.AsFloat64(),
			isValid: true,
		}
	case DivisibleNumber:
		tmp := v.(DivisibleNumber)
		if tmp.AsDouble().value == float64(0) {
			return DivisibleNumber{
				value:   1,
				isValid: false,
			}
		}
		return tmp
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for Double")
	return DivisibleNumber{
		value:   0,
		isValid: false,
	}
}

func (d *DivisibleNumber) AsInteger() Integer {
	return NewInteger(d.value)
}

func (d *DivisibleNumber) AsDouble() Double {
	return NewDouble(d.value)
}
