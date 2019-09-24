package types

import (
	"github.com/onsi/gomega"
	"math"
	"testing"
	"time"
)

func TestCreateNumberFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(int16(1)).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(int32(1)).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(int64(1)).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(float32(1.5)).AsFloat64()).
		To(gomega.Equal(1.5))
}

func TestCreateNumberFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsFloat64()).
		To(gomega.Equal(1.5))
}

func TestCreateNumberFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber("1.5").AsFloat64()).
		To(gomega.Equal(1.5))
	g.Expect(NewNumber("1").AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber([]byte("1.5")).AsFloat64()).
		To(gomega.Equal(1.5))
}

func TestCreateNumberFromNumberPointer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestCreateNumberFromNumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(*NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(1)))
}

func TestNumberNotPanicAndCreateDefaultNumberFromUnsupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(time.Now()).AsFloat64()).
		To(gomega.Equal(float64(0)))
	g.Expect(NewNumber(time.Now()).IsValid()).
		To(gomega.BeFalse())
}

func TestNoPanicAndDefaultNumberOnInvalidString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber("abc").AsFloat64()).
		To(gomega.Equal(float64(0)))
	g.Expect(NewNumber("abc").IsValid()).
		To(gomega.BeFalse())
}

func TestNoPanicAndDefaultNumberOnInvalidByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber([]byte("abc")).AsFloat64()).
		To(gomega.Equal(float64(0)))
	g.Expect(NewNumber([]byte("abc")).IsValid()).
		To(gomega.BeFalse())
}

func TestNumber_AsInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt()).
		To(gomega.Equal(1))
}

func TestNumber_AsInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt16()).
		To(gomega.Equal(int16(1)))
}

func TestNumber_AsInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt32()).
		To(gomega.Equal(int32(1)))
}

func TestNumber_AsInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt64()).
		To(gomega.Equal(int64(1)))
}

func TestNumber_AsFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsFloat32()).
		To(gomega.Equal(float32(1.5)))
}

func TestNumber_AsFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsFloat64()).
		To(gomega.Equal(1.5))
}

func TestNumber_AsString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsString(*NewNumber(2))).
		To(gomega.Equal("1.50"))
	g.Expect(NewNumber(1.5).AsString(*NewNumber(1))).
		To(gomega.Equal("1.5"))
	g.Expect(NewNumber(1.5).AsString(*NewNumber(0))).
		To(gomega.Equal("2"))
	g.Expect(NewNumber(1.5).AsString(*NewNumber(-1))).
		To(gomega.Equal("1.5"))
}

func TestNumber_AsByte(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(string(NewNumber(1.5).AsByte(*NewNumber(2)))).
		To(gomega.Equal("1.50"))
	g.Expect(string(NewNumber(1.5).AsByte(*NewNumber(1)))).
		To(gomega.Equal("1.5"))
	g.Expect(string(NewNumber(1.5).AsByte(*NewNumber(0)))).
		To(gomega.Equal("2"))
	g.Expect(string(NewNumber(1.5).AsByte(*NewNumber(-1)))).
		To(gomega.Equal("1.5"))
}

func TestNumber_IsInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []*Number{
		NewNumber(0),
		NewNumber(1),
		NewNumber(2),
		NewNumber(3),
	}
	g.Expect(z[1].IsInRange(*z[0], *z[2])).
		To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(*z[1], *z[2])).
		To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(*z[2], *z[3])).
		To(gomega.BeFalse())
}

func TestNumber_Equals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewNumber(1)
	x := NewNumber(2)
	g.Expect(z.Equals(*z)).
		To(gomega.BeTrue())
	g.Expect(z.Equals(*x)).
		To(gomega.BeFalse())
}

func TestNumber_IsAbove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []*Number{
		NewNumber(0),
		NewNumber(1),
	}
	g.Expect(z[0].IsAbove(*z[0])).
		To(gomega.BeFalse())
	g.Expect(z[0].IsAbove(*z[1])).
		To(gomega.BeFalse())
	g.Expect(z[1].IsAbove(*z[0])).
		To(gomega.BeTrue())
}

func TestNumber_IsBelow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []*Number{
		NewNumber(0),
		NewNumber(1),
	}
	g.Expect(z[0].IsBelow(*z[0])).
		To(gomega.BeFalse())
	g.Expect(z[0].IsBelow(*z[1])).
		To(gomega.BeTrue())
	g.Expect(z[1].IsBelow(*z[0])).
		To(gomega.BeFalse())
}

func TestNumber_Clamp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []*Number{
		NewNumber(100),
		NewNumber(50),
		NewNumber(200),
		NewNumber(400),
		NewNumber(1),
	}
	g.Expect(z[0].Clamp(*z[1], *z[2]).AsFloat64()).
		To(gomega.Equal(float64(100)))
	g.Expect(z[1].Clamp(*z[1], *z[2]).AsFloat64()).
		To(gomega.Equal(float64(50)))
	g.Expect(z[2].Clamp(*z[1], *z[2]).AsFloat64()).
		To(gomega.Equal(float64(200)))
	g.Expect(z[3].Clamp(*z[1], *z[2]).AsFloat64()).
		To(gomega.Equal(float64(200)))
	g.Expect(z[4].Clamp(*z[1], *z[2]).AsFloat64()).
		To(gomega.Equal(float64(50)))
}

func TestNumber_Add(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Add(*NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(6)))
}

func TestNumber_Subtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Subtract(*NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(4)))
}

func TestNumber_Multiply(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Multiply(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(float64(10)))
}

func TestNumber_Divide(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4).Divide(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(float64(2)))
}

func TestNumber_Divide_ByZero(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4).Divide(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(4)))
}

func TestNumber_Ceil(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4.006).Ceil(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(5)))
	g.Expect(NewNumber(6.004).Ceil(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(6.01))
}

func TestNumber_Floor(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4.006).Floor(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(4)))
	g.Expect(NewNumber(0.046).Floor(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(0.04))
}

func TestNumber_Round(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4.006).Round(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(4)))
	g.Expect(NewNumber(4.006).Round(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(4.01))
}

func TestRandomNumberInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	for i := 0; i < 1000; i++ {
		z := RandomNumberInt(*NewNumber(1), *NewNumber(10))
		g.Expect(z.AsFloat64() > 0 && z.AsFloat64() < 11).
			To(gomega.BeTrue())
	}
}

func TestRandomNumberFloat(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	for i := 0; i < 1000; i++ {
		z := RandomNumberFloat(*NewNumber(1.5), *NewNumber(10.5))
		g.Expect(z.AsFloat64() > 1.49 && z.AsFloat64() < 10.51).
			To(gomega.BeTrue())
	}
}

func TestNumber_Absolute(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Absolute().AsFloat64()).
		To(gomega.Equal(math.Abs(1.0)))
}

func TestNumber_ArcSine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ArcSine().AsFloat64()).
		To(gomega.Equal(math.Asin(1.0)))
}

func TestNumber_ArcCosine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ArcCosine().AsFloat64()).
		To(gomega.Equal(math.Acos(1.0)))
}

func TestNumber_ArcTangent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ArcTangent().AsFloat64()).
		To(gomega.Equal(math.Atan(1.0)))
}

func TestNumber_InverseHyperbolicSine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).InverseHyperbolicSine().AsFloat64()).
		To(gomega.Equal(math.Asinh(1.0)))
}

func TestNumber_InverseHyperbolicCosine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).InverseHyperbolicCosine().AsFloat64()).
		To(gomega.Equal(math.Acosh(1.0)))
}

func TestNumber_InverseHyperbolicTangent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).InverseHyperbolicTangent().AsFloat64()).
		To(gomega.Equal(math.Atanh(1.0)))
}

func TestNumber_ArcTangent2(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ArcTangent2(*NewNumber(5.2)).AsFloat64()).
		To(gomega.Equal(math.Atan2(1.0, 5.2)))
}

func TestNumber_Sine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Sine().AsFloat64()).
		To(gomega.Equal(math.Sin(1.0)))
}

func TestNumber_Cosine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Cosine().AsFloat64()).
		To(gomega.Equal(math.Cos(1.0)))
}

func TestNumber_Tangent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Tangent().AsFloat64()).
		To(gomega.Equal(math.Tan(1.0)))
}

func TestNumber_HyperbolicSine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).HyperbolicSine().AsFloat64()).
		To(gomega.Equal(math.Sinh(1.0)))
}

func TestNumber_HyperbolicCosine(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).HyperbolicCosine().AsFloat64()).
		To(gomega.Equal(math.Cosh(1.0)))
}

func TestNumber_HyperbolicTangent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).HyperbolicTangent().AsFloat64()).
		To(gomega.Equal(math.Tanh(1.0)))
}

func TestNumber_Base2Exponential(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Base2Exponential().AsFloat64()).
		To(gomega.Equal(math.Exp2(1.0)))
}

func TestNumber_BaseEExponential(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).BaseEExponential().AsFloat64()).
		To(gomega.Equal(math.Expm1(1.0)))
}

func TestNumber_BinaryExponent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).BinaryExponent().AsFloat64()).
		To(gomega.Equal(math.Logb(1.0)))
}

func TestNumber_BinaryExponential(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).BinaryExponential().AsInt()).
		To(gomega.Equal(math.Ilogb(1.0)))
}

func TestNumber_BinaryLogarithm(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).BinaryLogarithm().AsFloat64()).
		To(gomega.Equal(math.Log2(1.0)))
}

func TestNumber_ComplementaryErrorFunction(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ComplementaryErrorFunction().AsFloat64()).
		To(gomega.Equal(math.Erfc(1.0)))
}

func TestNumber_Copysign(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Copysign(*NewNumber(5.2)).AsFloat64()).
		To(gomega.Equal(math.Copysign(1.0, 5.2)))
}

func TestNumber_CubeRoot(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).CubeRoot().AsFloat64()).
		To(gomega.Equal(math.Cbrt(1.0)))
}

func TestNumber_DecimalLogarithm(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).DecimalLogarithm().AsFloat64()).
		To(gomega.Equal(math.Log10(1.0)))
}

func TestNumber_ErrorFunction(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).ErrorFunction().AsFloat64()).
		To(gomega.Equal(math.Erf(1.0)))
}

func TestNumber_Exponential(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Exponential().AsFloat64()).
		To(gomega.Equal(math.Exp(1.0)))
}

func TestNumber_Gamma(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Gamma().AsFloat64()).
		To(gomega.Equal(math.Gamma(1.0)))
}

func TestNumber_InverseComplementaryErrorFunction(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).InverseComplementaryErrorFunction().AsFloat64()).
		To(gomega.Equal(math.Erfcinv(1.0)))
}

func TestNumber_InverseErrorFunction(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).InverseErrorFunction().AsFloat64()).
		To(gomega.Equal(math.Erfinv(1.0)))
}

func TestNumber_IsValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).IsValid()).
		To(gomega.BeTrue())
}

func TestNumber_Modulo(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Modulo(*NewNumber(5.2)).AsFloat64()).
		To(gomega.Equal(math.Mod(1.0, 5.2)))
}

func TestNumber_NaturalLogarithm(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).NaturalLogarithm().AsFloat64()).
		To(gomega.Equal(math.Log(1.0)))
}

func TestNumber_NaturalLogarithmPlus1(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).NaturalLogarithmPlus1().AsFloat64()).
		To(gomega.Equal(math.Log1p(1.0)))
}

func TestNumber_OrderNBesselFirst(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderNBesselFirst(*NewNumber(5.2)).AsFloat64()).
		To(gomega.Equal(math.Jn(1.0, 5.2)))
}

func TestNumber_OrderNBesselSecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderNBesselSecond(*NewNumber(5.2)).AsFloat64()).
		To(gomega.Equal(math.Yn(1.0, 5.2)))
}

func TestNumber_OrderOneBesselFirst(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderOneBesselFirst().AsFloat64()).
		To(gomega.Equal(math.J1(1.0)))
}

func TestNumber_OrderOneBesselSecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderOneBesselSecond().AsFloat64()).
		To(gomega.Equal(math.Y1(1.0)))
}

func TestNumber_OrderZeroBesselFirst(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderZeroBesselFirst().AsFloat64()).
		To(gomega.Equal(math.J0(1.0)))
}

func TestNumber_OrderZeroBesselSecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).OrderZeroBesselSecond().AsFloat64()).
		To(gomega.Equal(math.Y0(1.0)))
}

func TestNumber_Power(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Power(*NewNumber(5)).AsFloat64()).
		To(gomega.Equal(math.Pow(1.0, 5)))
}

func TestNumber_Signbit(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).Signbit()).
		To(gomega.BeFalse())
}

func TestNumber_SquareRoot(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).SquareRoot().AsFloat64()).
		To(gomega.Equal(math.Sqrt(1.0)))
}
