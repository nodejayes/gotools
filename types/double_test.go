package types

import (
	"github.com/onsi/gomega"
	"testing"
	"time"
)

func TestCreateDoubleFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(1)
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(1)))
}

func TestCreateDoubleFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(int16(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(int32(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(int64(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(float32(1.5))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(1.5)
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble("1.5")
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble([]byte("1.5"))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(NewInteger(1))
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(1)))
}

func TestCreateDoubleFromDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(NewDouble(1.5))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestDoubleNotPanicAndCreateDefaultDoubleFromUnsupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(time.Now())
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(0)))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestNoPanicAndDefaultDoubleOnInvalidString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble("abc")
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(0)))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestNoPanicAndDefaultDoubleOnInvalidByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble([]byte("abc"))
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(0)))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestDoubleCanGetDoubleValueAsFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(1.5)
	g.Expect(i.AsFloat32()).To(gomega.Equal(float32(1.5)))
}

func TestDoubleCanGetDoubleValueAsFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(1.5)
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestDoubleCanGetAsString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := NewDouble(1.5)
	g.Expect(i.ToString(NewInteger(2))).To(gomega.Equal("1.50"))
	g.Expect(i.ToString(NewInteger(1))).To(gomega.Equal("1.5"))
	g.Expect(i.ToString(NewInteger(0))).To(gomega.Equal("2"))
	g.Expect(i.ToString(NewInteger(-1))).To(gomega.Equal("1.5"))
}

func TestDoubleIsInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Double{
		NewDouble(0),
		NewDouble(1),
		NewDouble(2),
		NewDouble(3),
	}
	g.Expect(z[1].IsInRange(z[0], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[1], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[2], z[3])).To(gomega.BeFalse())
}

func TestDoubleEquals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(1)
	x := NewDouble(2)
	g.Expect(z.Equals(z)).To(gomega.BeTrue())
	g.Expect(z.Equals(x)).To(gomega.BeFalse())
}

func TestDoubleIsAbove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Double{
		NewDouble(0),
		NewDouble(1),
	}
	g.Expect(z[0].IsAbove(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsAbove(z[1])).To(gomega.BeFalse())
	g.Expect(z[1].IsAbove(z[0])).To(gomega.BeTrue())
}

func TestDoubleIsBelow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Double{
		NewDouble(0),
		NewDouble(1),
	}
	g.Expect(z[0].IsBelow(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsBelow(z[1])).To(gomega.BeTrue())
	g.Expect(z[1].IsBelow(z[0])).To(gomega.BeFalse())
}

func TestDoubleClamp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []Double{
		NewDouble(100),
		NewDouble(50),
		NewDouble(200),
		NewDouble(400),
		NewDouble(1),
	}
	tmp1 := z[0].Clamp(z[1], z[2])
	tmp2 := z[1].Clamp(z[1], z[2])
	tmp3 := z[2].Clamp(z[1], z[2])
	tmp4 := z[3].Clamp(z[1], z[2])
	tmp5 := z[4].Clamp(z[1], z[2])
	g.Expect(tmp1.AsFloat64()).To(gomega.Equal(float64(100)))
	g.Expect(tmp2.AsFloat64()).To(gomega.Equal(float64(50)))
	g.Expect(tmp3.AsFloat64()).To(gomega.Equal(float64(200)))
	g.Expect(tmp4.AsFloat64()).To(gomega.Equal(float64(200)))
	g.Expect(tmp5.AsFloat64()).To(gomega.Equal(float64(50)))
}

func TestDoubleAdd(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(5)
	res := z.Add(NewDouble(1))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(6)))
}

func TestDoubleSubtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(5)
	res := z.Subtract(NewDouble(1))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))
}

func TestDoubleMultiply(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(5)
	res := z.Multiply(NewDouble(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(10)))
}

func TestDoubleDivide(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(4)
	res := z.Divide(NewDouble(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(2)))
}

func TestDoublePanicOnDivisionByZero(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewDouble(4)
	g.Expect(func() {
		z.Divide(NewDouble(0))
	}).Should(gomega.Panic())
}

func TestDoubleCeil(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDouble(4.006)
	res := example.Ceil(NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(5)))

	example = NewDouble(6.004)
	res = example.Ceil(NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(6.01))
}

func TestDoubleFloor(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDouble(4.006)
	res := example.Floor(NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))

	example = NewDouble(0.046)
	res = example.Floor(NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(0.04))
}

func TestDoubleRound(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDouble(4.006)
	res := example.Round(NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))

	example = NewDouble(4.006)
	res = example.Round(NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(4.01))
}

func TestRandomDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	min := NewDouble(1)
	max := NewDouble(10)
	z := RandomDouble(min, max)
	g.Expect(z.AsFloat64() > 0 && z.AsFloat64() < 11).To(gomega.BeTrue())
}
