package double_type

import (
	"github.com/nodejayes/gotools/src/types/integer"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

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
	case int_type.Integer:
		tmp := v.(int_type.Integer)
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
