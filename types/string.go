package types

import (
	"bytes"
	"reflect"
	"regexp"
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

// get the String as Byte Array
func (s *String) AsByte() []byte {
	return []byte(s.value)
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

// all letters are mapped to Upper Case
func (s *String) ToUpperCase() *String {
	return NewString(strings.ToUpper(s.value))
}

// all letters are mapped to Lower Case
func (s *String) ToLowerCase() *String {
	return NewString(strings.ToLower(s.value))
}

// check if the String equals the given String
func (s *String) Equal(v String) bool {
	return s.value == v.value
}

// add the given String behind this String
func (s *String) Concat(v String) *String {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(s.value)
	buf.WriteString(v.value)
	return NewString(buf.String())
}

// fill the String with the given template String to the given length
// on left and right side
func (s *String) Pad(length Number, template String) *String {
	myLength := s.Length()
	if myLength.Equals(length) || myLength.IsAbove(length) {
		return NewString(s)
	}
	var tmp []string
	var tmpBefore []string
	var tmpAfter []string
	sw := true
	for {
		for _, c := range template.AsString() {
			if !sw {
				tmpBefore = append(tmpBefore, string(c))
			} else {
				tmpAfter = append(tmpAfter, string(c))
			}
			returnStringLength := NewNumber((len(tmpBefore) + len(tmpAfter)) + s.Length().AsInt())
			if returnStringLength.Equals(length) || returnStringLength.IsAbove(length) {
				tmp = append(tmpBefore, s.value)
				tmp = append(tmp, tmpAfter...)
				return NewString(strings.Join(tmp, ""))
			}
		}
		sw = !sw
	}
}

// same as Pad but append the template on the left
func (s *String) PadLeft(length Number, template String) *String {
	return takeCharsByMaxLength(*s.Length(), length, template).Concat(*s)
}

// same as Pad but append the template on the right
func (s *String) PadRight(length Number, template String) *String {
	return s.Concat(*takeCharsByMaxLength(*s.Length(), length, template))
}

// check if the String begins with the Template
func (s *String) StartWith(template String) bool {
	idx := 0
	for _, c := range template.AsString() {
		if !s.CharAt(*NewNumber(idx + 1)).Equal(*NewString(string(c))) {
			return false
		}
		idx++
	}
	return true
}

// check if the String ends with the Template
func (s *String) EndWith(template String) bool {
	idx := s.Length().AsInt()
	for _, c := range template.Reverse().AsString() {
		c1 := *NewString(string(c))
		c2 := s.CharAt(*NewNumber(idx))
		if !c2.Equal(c1) {
			return false
		}
		idx--
	}
	return true
}

// turn around the String
func (s *String) Reverse() *String {
	runes := []rune(s.value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return NewString(string(runes))
}

// remove the template from the String at start and end
func (s *String) Trim(template String) *String {
	if !s.StartWith(template) && !s.EndWith(template) {
		return NewString(s.value)
	}
	return replaceStart(*replaceStart(*s, template).Reverse(), template).Reverse()
}

// remove the template from the String at start
func (s *String) TrimLeft(template String) *String {
	if !s.StartWith(template) {
		return NewString(s.value)
	}
	return replaceStart(*s, template)
}

// remove the template from the String at end
func (s *String) TrimRight(template String) *String {
	if !s.EndWith(template) {
		return NewString(s.value)
	}
	return replaceStart(*s.Reverse(), template).Reverse()
}

// repeat the String for the given times
func (s *String) Repeat(times Number) *String {
	tmp := EmptyString()
	for i := ZERO; i.IsBelow(times); i.Increment() {
		tmp = tmp.Concat(*s)
	}
	return tmp
}

// replace the first match in the String with the replacer
func (s *String) Replace(search String, replacer String) *String {
	return NewString(strings.Replace(s.value, search.value, replacer.value, 1))
}

// replace all matches in the String with the replacer
func (s *String) ReplaceAll(search String, replacer String) *String {
	return NewString(strings.Replace(s.value, search.value, replacer.value, -1))
}

// split the String by the template String
func (s *String) Split(template String) []*String {
	var tmp []*String
	for _, v := range strings.Split(s.value, template.value) {
		tmp = append(tmp, NewString(v))
	}
	return tmp
}

// insert the given template String at position
// if the position is outside the string
// the template was added before or behind
func (s *String) Insert(position Number, template String) *String {
	isBehind := position.IsAbove(*s.Length()) || position.Equals(*s.Length())
	isBefore := position.IsBelow(ZERO) || position.Equals(ZERO)
	if isBehind {
		return s.Concat(template)
	}
	if isBefore {
		return template.Concat(*s)
	}
	before := NewString(s.value[0:position.AsInt()])
	after := NewString(s.value[position.AsInt():])
	return before.Concat(template).Concat(*after)
}

// remove the count letters from position
func (s *String) Remove(position Number, count Number) *String {
	end := position.Add(count)
	noBefore := position.Equals(ZERO) || position.IsBelow(ZERO)
	noAfter := end.IsAbove(*s.Length())
	if noBefore && noAfter {
		return EmptyString()
	}
	before := EmptyString()
	after := EmptyString()
	if !noBefore {
		before = NewString(s.value[0:position.AsInt()])
	}
	if !noAfter {
		after = NewString(s.value[end.AsInt():])
	}
	return before.Concat(*after)
}

// get a String from position and for the Length of the second argument
func (s *String) SubString(position Number, length Number) *String {
	tmp := []rune(s.value)
	if length.IsBelow(ZERO) {
		return NewString(string(tmp[position.AsInt():]))
	}
	end := position.Add(length)
	return NewString(string(tmp[position.AsInt():end.AsInt()]))
}

// get the start Index of the template String from the First found match
func (s *String) IndexOf(template String) *Number {
	return NewNumber(strings.Index(s.value, template.value))
}

// get the start Index of the template String from the Last found match
func (s *String) LastIndexOf(template String) *Number {
	return NewNumber(strings.LastIndex(s.value, template.value))
}

// get the Text as String between the two Strings begin and end
func (s *String) TextBetween(begin, end String) []*String {
	var tmp []*String
	for _, split := range s.Split(begin) {
		subsplit := split.Split(end)
		if len(subsplit) < 1 || subsplit[0].Equal(*EmptyString()) {
			continue
		}
		tmp = append(tmp, subsplit[0])
	}
	return tmp
}

// check if the given search String in the String
func (s *String) Contains(search String) bool {
	return strings.Contains(s.value, search.value)
}

// count the matches of search in the String
func (s *String) ContainsCount(search String) *Number {
	return NewNumber(strings.Count(s.value, search.value))
}

// truncate the string by count with the replacer at the end
// the replacer was counted to the string length that was truncated!
func (s *String) Truncate(count Number, replacer String) *String {
	replacerIsToLong := replacer.Length().IsAbove(count) || replacer.Length().Equals(count)
	if replacerIsToLong {
		return NewString(replacer.value[0:count.AsInt()])
	}
	takeLength := count.Subtract(*replacer.Length())
	tmp := NewString(s.value[0:takeLength.AsInt()])
	return tmp.Concat(replacer)
}

// runs a Regular Expression again the String and return the result
func (s *String) RegularExpression(regularExpression String) bool {
	r, err := regexp.Match(regularExpression.value, s.AsByte())
	if err != nil {
		println("ERROR in String RegularExpression: " + err.Error())
		return false
	}
	return r
}

func takeCharsByMaxLength(stringLength, takenLength Number, template String) *String {
	if stringLength.Equals(takenLength) || stringLength.IsAbove(takenLength) {
		return EmptyString()
	}
	var tmp []string
	for {
		for _, c := range template.AsString() {
			tmp = append(tmp, string(c))
			returnStringLength := NewNumber((len(tmp)) + stringLength.AsInt())
			if returnStringLength.Equals(takenLength) || returnStringLength.IsAbove(takenLength) {
				return NewString(strings.Join(tmp, ""))
			}
		}
	}
}

func replaceStart(target, template String) *String {
	templateLength := template.Length().AsInt()
	tmp := target.Clone()
	for tmp.StartWith(template) {
		tmp = NewString(tmp.value[templateLength:])
	}
	return tmp
}
