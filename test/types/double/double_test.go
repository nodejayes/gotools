package double

import (
	"github.com/nodejayes/gotools/src/types/double"
	"github.com/nodejayes/gotools/src/types/integer"
	"github.com/onsi/gomega"
	"testing"
	"time"
)

func TestCreateDoubleFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(1)
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(1)))
}

func TestCreateDoubleFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(int16(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(int32(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(int64(1))
	g.Expect(i.AsFloat64(), gomega.Equal(1))
}

func TestCreateDoubleFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(float32(1.5))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(1.5)
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble("1.5")
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble([]byte("1.5"))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestCreateDoubleFromInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(int_type.NewInteger(1))
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(1)))
}

func TestCreateDoubleFromDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(double_type.NewDouble(1.5))
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestNotPanicAndCreateDefaultDoubleFromUnsupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(time.Now())
	g.Expect(i.AsFloat64()).To(gomega.Equal(float64(0)))
	g.Expect(i.IsValid()).To(gomega.BeFalse())
}

func TestCanGetDoubleValueAsFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(1.5)
	g.Expect(i.AsFloat32()).To(gomega.Equal(float32(1.5)))
}

func TestCanGetDoubleValueAsFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	i := double_type.NewDouble(1.5)
	g.Expect(i.AsFloat64()).To(gomega.Equal(1.5))
}

func TestIsInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []double_type.Double{
		double_type.NewDouble(0),
		double_type.NewDouble(1),
		double_type.NewDouble(2),
		double_type.NewDouble(3),
	}
	g.Expect(z[1].IsInRange(z[0], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[1], z[2])).To(gomega.BeTrue())
	g.Expect(z[1].IsInRange(z[2], z[3])).To(gomega.BeFalse())
}

func TestEquals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(1)
	x := double_type.NewDouble(2)
	g.Expect(z.Equals(z)).To(gomega.BeTrue())
	g.Expect(z.Equals(x)).To(gomega.BeFalse())
}

func TestIsAbove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []double_type.Double{
		double_type.NewDouble(0),
		double_type.NewDouble(1),
	}
	g.Expect(z[0].IsAbove(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsAbove(z[1])).To(gomega.BeFalse())
	g.Expect(z[1].IsAbove(z[0])).To(gomega.BeTrue())
}

func TestIsBelow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []double_type.Double{
		double_type.NewDouble(0),
		double_type.NewDouble(1),
	}
	g.Expect(z[0].IsBelow(z[0])).To(gomega.BeFalse())
	g.Expect(z[0].IsBelow(z[1])).To(gomega.BeTrue())
	g.Expect(z[1].IsBelow(z[0])).To(gomega.BeFalse())
}

func TestClamp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := []double_type.Double{
		double_type.NewDouble(100),
		double_type.NewDouble(50),
		double_type.NewDouble(200),
		double_type.NewDouble(400),
		double_type.NewDouble(1),
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

func TestAdd(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(5)
	res := z.Add(double_type.NewDouble(1))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(6)))
}

func TestSubtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(5)
	res := z.Subtract(double_type.NewDouble(1))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))
}

func TestMultiply(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(5)
	res := z.Multiply(double_type.NewDouble(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(10)))
}

func TestDivide(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(4)
	res := z.Divide(double_type.NewDouble(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(2)))
}

func TestPanicOnDivisionByZero(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	z := double_type.NewDouble(4)
	g.Expect(func() {
		z.Divide(double_type.NewDouble(0))
	}).Should(gomega.Panic())
}

func TestCeil(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := double_type.NewDouble(4.006)
	res := example.Ceil(int_type.NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(5)))

	example = double_type.NewDouble(6.004)
	res = example.Ceil(int_type.NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(6.01))
}

func TestFloor(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := double_type.NewDouble(4.006)
	res := example.Floor(int_type.NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))

	example = double_type.NewDouble(0.046)
	res = example.Floor(int_type.NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(0.04))
}

func TestRound(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := double_type.NewDouble(4.006)
	res := example.Round(int_type.NewInteger(0))
	g.Expect(res.AsFloat64()).To(gomega.Equal(float64(4)))

	example = double_type.NewDouble(4.006)
	res = example.Round(int_type.NewInteger(2))
	g.Expect(res.AsFloat64()).To(gomega.Equal(4.01))
}

func TestRandomDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	min := double_type.NewDouble(1)
	max := double_type.NewDouble(10)
	z := double_type.RandomDouble(min, max)
	g.Expect(z.AsFloat64() > 0 && z.AsFloat64() < 11).To(gomega.BeTrue())
}
