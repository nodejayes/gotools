package int_type

import (
	"github.com/nodejayes/gotools/src/types/double"
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
	case double_type.Double:
		tmp := v.(double_type.Double)
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
