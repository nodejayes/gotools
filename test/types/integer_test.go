package types

import (
	"github.com/nodejayes/tooling/src/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestInteger(t *testing.T) {
	Convey("Integer", t, func() {
		globalI := types.NewInteger(4)
		Convey("can create new Integer Values from int", func() {
			i := types.NewInteger(1)
			So(i.AsInt(), ShouldEqual, 1)
			So(i.IsValid(), ShouldBeTrue)
		})
		Convey("can create new Integer Values from int16", func() {
			i := types.NewInteger(int16(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from int32", func() {
			i := types.NewInteger(int32(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from int64", func() {
			i := types.NewInteger(int64(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from float32", func() {
			i := types.NewInteger(float32(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from float64", func() {
			i := types.NewInteger(float64(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from string", func() {
			i := types.NewInteger("1")
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from []byte", func() {
			i := types.NewInteger([]byte("1"))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from Integer", func() {
			i := types.NewInteger(types.NewInteger(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("can create new Integer Values from Double", func() {
			i := types.NewInteger(types.NewDouble(1))
			So(i.AsInt(), ShouldEqual, 1)
		})
		Convey("log info when type not supported", func() {
			i := types.NewInteger(time.Now())
			So(i.AsInt(), ShouldEqual, 0)
			So(i.IsValid(), ShouldBeFalse)
		})
		Convey("float values are round", func() {
			i := types.NewInteger(float32(1.5))
			So(i.AsInt(), ShouldEqual, 1)
			i = types.NewInteger(float32(1.4))
			So(i.AsInt(), ShouldEqual, 1)
			i = types.NewInteger(float32(1.6))
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
			z := []types.Integer{
				types.NewInteger(0),
				types.NewInteger(1),
				types.NewInteger(2),
				types.NewInteger(3),
			}
			So(z[1].IsInRange(z[0], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[1], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[2], z[3]), ShouldBeFalse)
		})
		Convey("check Equals", func() {
			z := types.NewInteger(1)
			x := types.NewInteger(2)
			So(z.Equals(z), ShouldBeTrue)
			So(z.Equals(x), ShouldBeFalse)
		})
		Convey("check IsAbove", func() {
			z := []types.Integer{
				types.NewInteger(0),
				types.NewInteger(1),
			}
			So(z[0].IsAbove(z[0]), ShouldBeFalse)
			So(z[0].IsAbove(z[1]), ShouldBeFalse)
			So(z[1].IsAbove(z[0]), ShouldBeTrue)
		})
		Convey("check IsBelow", func() {
			z := []types.Integer{
				types.NewInteger(0),
				types.NewInteger(1),
			}
			So(z[0].IsBelow(z[0]), ShouldBeFalse)
			So(z[0].IsBelow(z[1]), ShouldBeTrue)
			So(z[1].IsBelow(z[0]), ShouldBeFalse)
		})
		Convey("check Clamp", func() {
			z := []types.Integer{
				types.NewInteger(100),
				types.NewInteger(50),
				types.NewInteger(200),
				types.NewInteger(400),
				types.NewInteger(1),
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
			z := types.NewInteger(5)
			res := z.Add(types.NewInteger(1))
			So(res.AsInt(), ShouldEqual, 6)
		})
		Convey("check Subtract", func() {
			z := types.NewInteger(5)
			res := z.Subtract(types.NewInteger(1))
			So(res.AsInt(), ShouldEqual, 4)
		})
		Convey("check Multiply", func() {
			z := types.NewInteger(5)
			res := z.Multiply(types.NewInteger(2))
			So(res.AsInt(), ShouldEqual, 10)
		})
		Convey("check Division", func() {
			z := types.NewInteger(4)
			res := z.Divide(types.NewInteger(2))
			So(res.AsInt(), ShouldEqual, 2)
		})
		Convey("no Division by Zero", func() {
			z := types.NewInteger(4)
			So(func() {
				z.Divide(types.NewInteger(0))
			}, ShouldPanic)
		})
		Convey("can create Random Integer", func() {
			min := types.NewInteger(1)
			max := types.NewInteger(10)
			z := types.RandomInteger(min, max)
			So(z.AsInt(), ShouldBeBetween, 0, 11)
		})
	})
}
