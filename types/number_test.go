package types

import (
	"github.com/onsi/gomega"
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

func TestNumberCanGetNumberValueAsInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt()).
		To(gomega.Equal(1))
}

func TestNumberCanGetNumberValueAsInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt16()).
		To(gomega.Equal(int16(1)))
}

func TestNumberCanGetNumberValueAsInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt32()).
		To(gomega.Equal(int32(1)))
}

func TestNumberCanGetNumberValueAsInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1).AsInt64()).
		To(gomega.Equal(int64(1)))
}

func TestNumberCanGetNumberValueAsFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsFloat32()).
		To(gomega.Equal(float32(1.5)))
}

func TestNumberCanGetNumberValueAsFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(1.5).AsFloat64()).
		To(gomega.Equal(1.5))
}

func TestNumberCanGetAsString(t *testing.T) {
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

func TestNumberCanGetAsByteArray(t *testing.T) {
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

func TestNumberIsInRange(t *testing.T) {
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

func TestNumberEquals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := NewNumber(1)
	x := NewNumber(2)
	g.Expect(z.Equals(*z)).
		To(gomega.BeTrue())
	g.Expect(z.Equals(*x)).
		To(gomega.BeFalse())
}

func TestNumberIsAbove(t *testing.T) {
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

func TestNumberIsBelow(t *testing.T) {
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

func TestNumberClamp(t *testing.T) {
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

func TestNumberAdd(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Add(*NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(6)))
}

func TestNumberSubtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Subtract(*NewNumber(1)).AsFloat64()).
		To(gomega.Equal(float64(4)))
}

func TestNumberMultiply(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(5).Multiply(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(float64(10)))
}

func TestNumberDivide(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4).Divide(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(float64(2)))
}

func TestNumberDivisionByZero(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4).Divide(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(4)))
}

func TestNumberCeil(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4.006).Ceil(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(5)))
	g.Expect(NewNumber(6.004).Ceil(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(6.01))
}

func TestNumberFloor(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewNumber(4.006).Floor(*NewNumber(0)).AsFloat64()).
		To(gomega.Equal(float64(4)))
	g.Expect(NewNumber(0.046).Floor(*NewNumber(2)).AsFloat64()).
		To(gomega.Equal(0.04))
}

func TestNumberRound(t *testing.T) {
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
