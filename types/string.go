package types

import (
	"reflect"
	"strings"
)

// represent a String Value
type String struct {
	value   string
	isValid bool
}

// create a new String
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
	case uint8:
		return &String{
			value:   string(v.(uint8)),
			isValid: true,
		}
	case Number:
		tmp := v.(Number)
		return &String{
			value:   tmp.AsString(*NewNumber(-1)),
			isValid: true,
		}
	case *Number:
		return &String{
			value:   v.(*Number).AsString(*NewNumber(-1)),
			isValid: true,
		}
	case String:
		tmp := v.(String)
		return &tmp
	case *String:
		tmp := v.(*String)
		return tmp
	}
	println("type " + reflect.TypeOf(v).Name() + " not supported for String")
	return EmptyString()
}

// returns a EmptyString value
func EmptyString() *String {
	return &String{
		value:   "",
		isValid: false,
	}
}

// is this String a Valid Value
func (s *String) IsValid() bool {
	return s.isValid
}

// Clone the Instance of this String into a new one
func (s *String) Clone() *String {
	tmp := NewString(*s)
	return tmp
}

// get the String as a Go String
func (s *String) AsString() string {
	return s.value
}

// get the String as a Number
func (s *String) AsNumber() *Number {
	return NewNumber(s.value)
}

// get the num of Characters in the String
func (s *String) Length() *Number {
	return NewNumber(len(s.value))
}

// get a Character as String Instance at a Position on the String
// position is the Index of the String Array +1
// to get the first Character use position = 1 not 0!
func (s *String) CharAt(position Number) *String {
	stringLength := *NewNumber(len(s.value))
	if position.IsAbove(stringLength) || position.IsBelow(ZERO) || position.Equals(ZERO) {
		return EmptyString()
	}
	index := position.Subtract(*NewNumber(1)).AsInt()
	c := s.value[index]
	return NewString(c)
}

func (s *String) ToUpperCase() *String {
	return NewString(strings.ToUpper(s.value))
}

func (s *String) ToLowerCase() *String {
	return NewString(strings.ToLower(s.value))
}

func (s *String) Equal(v String) bool {
	return s.value == v.value
}

func (s *String) Pad(length Number, template String) *String {
	return EmptyString()
}

func (s *String) PadLeft(length Number, template String) *String {
	return EmptyString()
}

func (s *String) PadRight(length Number, template String) *String {
	return EmptyString()
}

func (s *String) Trim(template String) *String {
	return EmptyString()
}

func (s *String) TrimLeft(template String) *String {
	return EmptyString()
}

func (s *String) TrimRight(template String) *String {
	return EmptyString()
}

func (s *String) Repeat(times Number) *String {
	return EmptyString()
}

func (s *String) Replace(search String, replacer String) *String {
	return EmptyString()
}

func (s *String) ReplaceAll(search String, replacer String) *String {
	return EmptyString()
}

func (s *String) Split(template String) *String {
	return EmptyString()
}

func (s *String) Insert(position Number, template String) *String {
	return EmptyString()
}

func (s *String) Remove(position Number, count Number) *String {
	return EmptyString()
}

func (s *String) SubString(position Number, length Number) *String {
	return EmptyString()
}

func (s *String) IndexOf(template String) *Number {
	return NewNumber(0)
}

func (s *String) LastIndexOf(template String) *Number {
	return NewNumber(0)
}

func (s *String) TextBetween(begin, end String) []*String {
	return []*String{EmptyString()}
}
