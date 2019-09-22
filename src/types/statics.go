package types

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

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

// Generates a new Integer random number in the range of min and max
func RandomInteger(min, max Integer) Integer {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewInteger(rand.Intn(max.AsInt()-min.AsInt()) + min.AsInt())
}

// Generates a new Double random number in the range of min and max
func RandomDouble(min, max Double) Double {
	rand.Seed(time.Now().UTC().UnixNano())
	return NewDouble(min.AsFloat64() + rand.Float64()*(max.AsFloat64()-min.AsFloat64()))
}
