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

var ZERO = *NewNumber(0)

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

// clone the Number into a new Instance
func (i *Number) Clone() *Number {
	tmp := NewNumber(*i)
	return tmp
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

// calculate the absolute value of the current Number
func (i *Number) Absolute() *Number {
	return NewNumber(math.Abs(i.value))
}

// get the ArcSine from the current Number
func (i *Number) ArcSine() *Number {
	return NewNumber(math.Asin(i.value))
}

// get the ArcCosine from the current Number
func (i *Number) ArcCosine() *Number {
	return NewNumber(math.Acos(i.value))
}

// get the ArcTangent from the current Number
func (i *Number) ArcTangent() *Number {
	return NewNumber(math.Atan(i.value))
}

func (i *Number) InverseHyperbolicSine() *Number {
	return NewNumber(math.Asinh(i.value))
}

func (i *Number) InverseHyperbolicCosine() *Number {
	return NewNumber(math.Acosh(i.value))
}

func (i *Number) InverseHyperbolicTangent() *Number {
	return NewNumber(math.Atanh(i.value))
}

func (i *Number) ArcTangent2(y Number) *Number {
	return NewNumber(math.Atan2(i.value, y.AsFloat64()))
}

func (i *Number) Cosine() *Number {
	return NewNumber(math.Cos(i.value))
}

func (i *Number) Sine() *Number {
	return NewNumber(math.Sin(i.value))
}

func (i *Number) Tangent() *Number {
	return NewNumber(math.Tan(i.value))
}

func (i *Number) HyperbolicSine() *Number {
	return NewNumber(math.Sinh(i.value))
}

func (i *Number) HyperbolicCosine() *Number {
	return NewNumber(math.Cosh(i.value))
}

func (i *Number) HyperbolicTangent() *Number {
	return NewNumber(math.Tanh(i.value))
}

func (i *Number) CubeRoot() *Number {
	return NewNumber(math.Cbrt(i.value))
}

func (i *Number) Copysign(v Number) *Number {
	return NewNumber(math.Copysign(i.value, v.AsFloat64()))
}

func (i *Number) ErrorFunction() *Number {
	return NewNumber(math.Erf(i.value))
}

func (i *Number) ComplementaryErrorFunction() *Number {
	return NewNumber(math.Erfc(i.value))
}

func (i *Number) InverseErrorFunction() *Number {
	return NewNumber(math.Erfinv(i.value))
}

func (i *Number) InverseComplementaryErrorFunction() *Number {
	return NewNumber(math.Erfcinv(i.value))
}

func (i *Number) Exponential() *Number {
	return NewNumber(math.Exp(i.value))
}

func (i *Number) Base2Exponential() *Number {
	return NewNumber(math.Exp2(i.value))
}

func (i *Number) BaseEExponential() *Number {
	return NewNumber(math.Expm1(i.value))
}

func (i *Number) Gamma() *Number {
	return NewNumber(math.Gamma(i.value))
}

func (i *Number) BinaryExponential() *Number {
	return NewNumber(math.Ilogb(i.value))
}

func (i *Number) SquareRoot() *Number {
	return NewNumber(math.Sqrt(i.value))
}

func (i *Number) Power(n Number) *Number {
	return NewNumber(math.Pow(i.value, n.AsFloat64()))
}

func (i *Number) NaturalLogarithm() *Number {
	return NewNumber(math.Log(i.value))
}

func (i *Number) NaturalLogarithmPlus1() *Number {
	return NewNumber(math.Log1p(i.value))
}

func (i *Number) BinaryLogarithm() *Number {
	return NewNumber(math.Log2(i.value))
}

func (i *Number) DecimalLogarithm() *Number {
	return NewNumber(math.Log10(i.value))
}

func (i *Number) BinaryExponent() *Number {
	return NewNumber(math.Logb(i.value))
}

func (i *Number) Signbit() bool {
	return math.Signbit(i.value)
}

func (i *Number) Modulo(n Number) *Number {
	return NewNumber(math.Mod(i.value, n.AsFloat64()))
}

func (i *Number) OrderZeroBesselFirst() *Number {
	return NewNumber(math.J0(i.value))
}

func (i *Number) OrderOneBesselFirst() *Number {
	return NewNumber(math.J1(i.value))
}

func (i *Number) OrderNBesselFirst(n Number) *Number {
	return NewNumber(math.Jn(i.AsInt(), n.AsFloat64()))
}

func (i *Number) OrderZeroBesselSecond() *Number {
	return NewNumber(math.Y0(i.value))
}

func (i *Number) OrderOneBesselSecond() *Number {
	return NewNumber(math.Y1(i.value))
}

func (i *Number) OrderNBesselSecond(n Number) *Number {
	return NewNumber(math.Yn(i.AsInt(), n.AsFloat64()))
}

// increment the Number this means it was mutated!
func (i *Number) Increment() {
	i.value++
}

// decrement the Number this means it was mutated!
func (i *Number) Decrement() {
	i.value--
}

// check if the Number is in a List of Numbers
func (i *Number) IsIn(list []*Number) bool {
	for _, n := range list {
		if i.Equals(*n) {
			return true
		}
	}
	return false
}
