package types

import "reflect"

type String struct {
	value   string
	isValid bool
}

func NewString(v interface{}) *String {
	switch v.(type) {
	case string:
		return &String{
			value:   v.(string),
			isValid: true,
		}
	case []byte:
		return &String{
			value:   string(v.([]byte)),
			isValid: true,
		}
	case int, int16, int32, int64:
		return &String{
			value:   NewNumber(v).AsString(*NewNumber(0)),
			isValid: true,
		}
	case float32, float64:
		return &String{
			value:   NewNumber(v).AsString(*NewNumber(-1)),
			isValid: true,
		}
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for String")
	return &String{
		value:   "",
		isValid: false,
	}
}

func (s *String) AsString() string {
	return s.value
}

func (s *String) AsNumber() *Number {
	return NewNumber(s.value)
}
