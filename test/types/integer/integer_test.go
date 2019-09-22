package integer

import (
	"github.com/nodejayes/gotools/src/types/double"
	"github.com/nodejayes/gotools/src/types/integer"
	"github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateNewIntegerFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCreateNewIntegerFromDouble(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestNoPanicAndDefaultIntegerWhenNotSupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestRoundingFloatValues(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCanGetIntegerValueAsInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCanGetIntegerValueAsInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCanGetIntegerValueAsInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestCanGetIntegerValueAsInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestIsInRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
}

func TestInteger(t *testing.T) {
	Convey("Integer", t, func() {
		globalI := int_type.NewInteger(4)
		Convey("can create new Integer Values from int", func() {
			i := int_type.NewInteger(1)
			So(i.AsInt(), ShouldEqual, 1)
			So(i.IsValid(), ShouldBeTrue)
		})
		Convey("can create new Integer Values from int16", func() {
			i := int_type.NewInteger(int16(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from int32", func() {
			i := int_type.NewInteger(int32(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from int64", func() {
			i := int_type.NewInteger(int64(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from float32", func() {
			i := int_type.NewInteger(float32(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from float64", func() {
			i := int_type.NewInteger(float64(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from string", func() {
			i := int_type.NewInteger("1")
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from []byte", func() {
			i := int_type.NewInteger([]byte("1"))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from Integer", func() {
			i := int_type.NewInteger(int_type.NewInteger(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from Double", func() {
			i := int_type.NewInteger(double_type.NewDouble(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("log info when type not supported", func() {
			i := int_type.NewInteger(time.Now())
			So(i.AsInt(), ShouldEqual, 0)
			So(i.IsValid(), ShouldBeFalse)
		})
		Convey("float values are round", func() {
			i := int_type.NewInteger(float32(1.5))
			So(i.AsInt(), ShouldEqual, 1)
			i = int_type.NewInteger(float32(1.4))
			So(i.AsInt(), ShouldEqual, 1)
			i = int_type.NewInteger(float32(1.6))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can get as int", func() {
			So(globalI.AsInt(), ShouldEqual, 4)
		})
		Convey("can get as int16", func() {
			So(globalI.AsInt16(), ShouldEqual, int16(4))
		})
		Convey("can get as int32", func() {
			So(globalI.AsInt32(), ShouldEqual, int32(4))
		})
		Convey("can get as int64", func() {
			So(globalI.AsInt64(), ShouldEqual, int64(4))
		})
		Convey("check IsInRange", func() {
			z := []int_type.Integer{
				int_type.NewInteger(0),
				int_type.NewInteger(1),
				int_type.NewInteger(2),
				int_type.NewInteger(3),
			}
			So(z[1].IsInRange(z[0], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[1], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[2], z[3]), ShouldBeFalse)
		})
		Convey("check Equals", func() {
			z := int_type.NewInteger(1)
			x := int_type.NewInteger(2)
			So(z.Equals(z), ShouldBeTrue)
			So(z.Equals(x), ShouldBeFalse)
		})
		Convey("check IsAbove", func() {
			z := []int_type.Integer{
				int_type.NewInteger(0),
				int_type.NewInteger(1),
			}
			So(z[0].IsAbove(z[0]), ShouldBeFalse)
			So(z[0].IsAbove(z[1]), ShouldBeFalse)
			So(z[1].IsAbove(z[0]), ShouldBeTrue)
		})
		Convey("check IsBelow", func() {
			z := []int_type.Integer{
				int_type.NewInteger(0),
				int_type.NewInteger(1),
			}
			So(z[0].IsBelow(z[0]), ShouldBeFalse)
			So(z[0].IsBelow(z[1]), ShouldBeTrue)
			So(z[1].IsBelow(z[0]), ShouldBeFalse)
		})
		Convey("check Clamp", func() {
			z := []int_type.Integer{
				int_type.NewInteger(100),
				int_type.NewInteger(50),
				int_type.NewInteger(200),
				int_type.NewInteger(400),
				int_type.NewInteger(1),
			}
			tmp1 := z[0].Clamp(z[1], z[2])
			tmp2 := z[1].Clamp(z[1], z[2])
			tmp3 := z[2].Clamp(z[1], z[2])
			tmp4 := z[3].Clamp(z[1], z[2])
			tmp5 := z[4].Clamp(z[1], z[2])
			So(tmp1.AsInt(), ShouldEqual, 100)
			So(tmp2.AsInt(), ShouldEqual, 50)
			So(tmp3.AsInt(), ShouldEqual, 200)
			So(tmp4.AsInt(), ShouldEqual, 200)
			So(tmp5.AsInt(), ShouldEqual, 50)
		})
		Convey("check Add", func() {
			z := int_type.NewInteger(5)
			res := z.Add(int_type.NewInteger(1))
			So(res.AsInt(), ShouldEqual, 6)
		})
		Convey("check Subtract", func() {
			z := int_type.NewInteger(5)
			res := z.Subtract(int_type.NewInteger(1))
			So(res.AsInt(), ShouldEqual, 4)
		})
		Convey("check Multiply", func() {
			z := int_type.NewInteger(5)
			res := z.Multiply(int_type.NewInteger(2))
			So(res.AsInt(), ShouldEqual, 10)
		})
		Convey("check Division", func() {
			z := int_type.NewInteger(4)
			res := z.Divide(int_type.NewInteger(2))
			So(res.AsInt(), ShouldEqual, 2)
		})
		Convey("no Division by Zero", func() {
			z := int_type.NewInteger(4)
			So(func() {
				z.Divide(int_type.NewInteger(0))
			}, ShouldPanic)
		})
		Convey("can create Random Integer", func() {
			min := int_type.NewInteger(1)
			max := int_type.NewInteger(10)
			z := int_type.RandomInteger(min, max)
			So(z.AsInt(), ShouldBeBetween, 0, 11)
		})
	})
}
