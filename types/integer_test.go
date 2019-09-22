package types

import (
	"github.com/onsi/gomega"
	"testing"
	"time"
)

func TestCreateNewIntegerFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1)
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(int16(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(int32(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(int64(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(float32(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(float64(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger("1")
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger([]byte("1"))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(NewInteger(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestCreateNewIntegerFromDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(NewDouble(1))
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestNoPanicAndDefaultIntegerWhenNotSupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(time.Now())
	g.Expect(i.AsInt()).To(gomega.Equal(0))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestNoPanicAndDefaultIntegerOnInvalidString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger("abc")
	g.Expect(i.AsInt()).To(gomega.Equal(0))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestNoPanicAndDefaultIntegerOnInvalidByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger([]byte("abc"))
	g.Expect(i.AsInt()).To(gomega.Equal(0))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestIntegerRoundingFloatValues(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1.5)
	g.Expect(i.AsInt()).To(gomega.Equal(1))
	g.Expect(i.IsValid()).To(gomega.BeTrue())
}

func TestIntegerCanGetIntegerValueAsInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1)
	g.Expect(i.AsInt()).To(gomega.Equal(1))
}

func TestIntegerCanGetIntegerValueAsInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1)
	g.Expect(i.AsInt16()).To(gomega.Equal(int16(1)))
}

func TestIntegerCanGetIntegerValueAsInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1)
	g.Expect(i.AsInt32()).To(gomega.Equal(int32(1)))
}

func TestIntegerCanGetIntegerValueAsInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(1)
	g.Expect(i.AsInt64()).To(gomega.Equal(int64(1)))
}

func TestIntegerCanGetAsString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewInteger(5)
	g.Expect(i.ToString()).To(gomega.Equal("5"))
}

func TestIntegerIsInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Integer{
		NewInteger(0),
		NewInteger(1),
		NewInteger(2),
		NewInteger(3),
	}
	g.Expect(z[1].IsInRange(z[0], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[1], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[2], z[3])).To(gomega.BeFalse())
}

func TestIntegerEquals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(1)
	x := NewInteger(2)
	g.Expect(z.Equals(z)).To(gomega.BeTrue())
	g.Expect(z.Equals(x)).To(gomega.BeFalse())
}

func TestIntegerIsAbove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Integer{
		NewInteger(0),
		NewInteger(1),
	}
	g.Expect(z[0].IsAbove(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsAbove(z[1])).To(gomega.BeFalse())
	g.Expect(z[1].IsAbove(z[0])).To(gomega.BeTrue())
}

func TestIntegerIsBelow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Integer{
		NewInteger(0),
		NewInteger(1),
	}
	g.Expect(z[0].IsBelow(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsBelow(z[1])).To(gomega.BeTrue())
	g.Expect(z[1].IsBelow(z[0])).To(gomega.BeFalse())
}

func TestIntegerClamp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Integer{
		NewInteger(100),
		NewInteger(50),
		NewInteger(200),
		NewInteger(400),
		NewInteger(1),
	}
	tmp1 := z[0].Clamp(z[1], z[2])
	tmp2 := z[1].Clamp(z[1], z[2])
	tmp3 := z[2].Clamp(z[1], z[2])
	tmp4 := z[3].Clamp(z[1], z[2])
	tmp5 := z[4].Clamp(z[1], z[2])
	g.Expect(tmp1.AsInt()).To(gomega.Equal(100))
	g.Expect(tmp2.AsInt()).To(gomega.Equal(50))
	g.Expect(tmp3.AsInt()).To(gomega.Equal(200))
	g.Expect(tmp4.AsInt()).To(gomega.Equal(200))
	g.Expect(tmp5.AsInt()).To(gomega.Equal(50))
}

func TestIntegerAdd(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(5)
	res := z.Add(NewInteger(1))
	g.Expect(res.AsInt()).To(gomega.Equal(6))
}

func TestIntegerSubtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(5)
	res := z.Subtract(NewInteger(1))
	g.Expect(res.AsInt()).To(gomega.Equal(4))
}

func TestIntegerMultiply(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(5)
	res := z.Multiply(NewInteger(2))
	g.Expect(res.AsInt()).To(gomega.Equal(10))
}

func TestIntegerDivide(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(4)
	res := z.Divide(NewInteger(2))
	g.Expect(res.AsInt()).To(gomega.Equal(2))
}

func TestIntegerPanicOnDivisionByZero(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewInteger(4)
	g.Expect(func() {
		z.Divide(NewInteger(0))
	}).To(gomega.Panic())
}

func TestRandomInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	min := NewInteger(1)
	max := NewInteger(10)
	z := RandomInteger(min, max)
	g.Expect(z.AsInt() > 0 && z.AsInt() < 11).To(gomega.BeTrue())
}
