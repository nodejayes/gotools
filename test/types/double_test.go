package types

import (
	"github.com/nodejayes/gotools/src/types"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestDouble(t *testing.T) {
	Convey("Double", t, func() {
		Convey("can create new Double Values from int", func() {
			i := types.NewDouble(1)
			So(i.AsFloat64(), ShouldEqual, 1)
		})
		Convey("can create new Double Values from int16", func() {
			i := types.NewDouble(int16(1))
			So(i.AsFloat64(), ShouldEqual, 1)
		})
		Convey("can create new Double Values from int32", func() {
			i := types.NewDouble(int32(1))
			So(i.AsFloat64(), ShouldEqual, 1)
		})
		Convey("can create new Double Values from int64", func() {
			i := types.NewDouble(int64(1))
			So(i.AsFloat64(), ShouldEqual, 1)
		})
		Convey("can create new Double Values from float32", func() {
			i := types.NewDouble(float32(1.5))
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("can create new Double Values from float64", func() {
			i := types.NewDouble(1.5)
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("can create new Double Values from String", func() {
			i := types.NewDouble("1.5")
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("can create new Double Values from []byte", func() {
			i := types.NewDouble([]byte("1.5"))
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("can create new Double Values from Integer", func() {
			i := types.NewDouble(types.NewInteger(1))
			So(i.AsFloat64(), ShouldEqual, 1)
		})
		Convey("can create new Double Values from Double", func() {
			i := types.NewDouble(types.NewDouble(1.5))
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("can get as float32", func() {
			i := types.NewDouble(1.5)
			So(i.AsFloat32(), ShouldEqual, float32(1.5))
		})
		Convey("can get as float64", func() {
			i := types.NewDouble(1.5)
			So(i.AsFloat64(), ShouldEqual, 1.5)
		})
		Convey("log info when type not supported", func() {
			i := types.NewDouble(time.Now())
			So(i.AsFloat64(), ShouldEqual, 0)
			So(i.IsValid(), ShouldBeFalse)
		})
		Convey("check IsInRange", func() {
			z := []types.Double{
				types.NewDouble(0),
				types.NewDouble(1),
				types.NewDouble(2),
				types.NewDouble(3),
			}
			So(z[1].IsInRange(z[0], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[1], z[2]), ShouldBeTrue)
			So(z[1].IsInRange(z[2], z[3]), ShouldBeFalse)
		})
		Convey("check Equals", func() {
			z := types.NewDouble(1)
			x := types.NewDouble(2)
			So(z.Equals(z), ShouldBeTrue)
			So(z.Equals(x), ShouldBeFalse)
		})
		Convey("check IsAbove", func() {
			z := []types.Double{
				types.NewDouble(0),
				types.NewDouble(1),
			}
			So(z[0].IsAbove(z[0]), ShouldBeFalse)
			So(z[0].IsAbove(z[1]), ShouldBeFalse)
			So(z[1].IsAbove(z[0]), ShouldBeTrue)
		})
		Convey("check IsBelow", func() {
			z := []types.Double{
				types.NewDouble(0),
				types.NewDouble(1),
			}
			So(z[0].IsBelow(z[0]), ShouldBeFalse)
			So(z[0].IsBelow(z[1]), ShouldBeTrue)
			So(z[1].IsBelow(z[0]), ShouldBeFalse)
		})
		Convey("check Clamp", func() {
			z := []types.Double{
				types.NewDouble(100),
				types.NewDouble(50),
				types.NewDouble(200),
				types.NewDouble(400),
				types.NewDouble(1),
			}
			tmp1 := z[0].Clamp(z[1], z[2])
			tmp2 := z[1].Clamp(z[1], z[2])
			tmp3 := z[2].Clamp(z[1], z[2])
			tmp4 := z[3].Clamp(z[1], z[2])
			tmp5 := z[4].Clamp(z[1], z[2])
			So(tmp1.AsFloat64(), ShouldEqual, 100)
			So(tmp2.AsFloat64(), ShouldEqual, 50)
			So(tmp3.AsFloat64(), ShouldEqual, 200)
			So(tmp4.AsFloat64(), ShouldEqual, 200)
			So(tmp5.AsFloat64(), ShouldEqual, 50)
		})
		Convey("check Add", func() {
			z := types.NewDouble(5)
			res := z.Add(types.NewDouble(1))
			So(res.AsFloat64(), ShouldEqual, 6)
		})
		Convey("check Subtract", func() {
			z := types.NewDouble(5)
			res := z.Subtract(types.NewDouble(1))
			So(res.AsFloat64(), ShouldEqual, 4)
		})
		Convey("check Multiply", func() {
			z := types.NewDouble(5)
			res := z.Multiply(types.NewDouble(2))
			So(res.AsFloat64(), ShouldEqual, 10)
		})
		Convey("check Division", func() {
			z := types.NewDouble(4)
			res := z.Divide(types.NewDouble(2))
			So(res.AsFloat64(), ShouldEqual, 2)
		})
		Convey("no Division by Zero", func() {
			z := types.NewDouble(4)
			So(func() {
				z.Divide(types.NewDouble(0))
			}, ShouldPanic)
		})
		Convey("can Ceil Double", func() {
			example := types.NewDouble(4.006)
			res := example.Ceil(types.NewInteger(0))
			So(res.AsFloat64(), ShouldEqual, 5)
			example = types.NewDouble(6.004)
			res = example.Ceil(types.NewInteger(2))
			So(res.AsFloat64(), ShouldEqual, 6.01)
		})
		Convey("can Floor Double", func() {
			example := types.NewDouble(4.006)
			res := example.Floor(types.NewInteger(0))
			So(res.AsFloat64(), ShouldEqual, 4)
			example = types.NewDouble(0.046)
			res = example.Floor(types.NewInteger(2))
			So(res.AsFloat64(), ShouldEqual, 0.04)
		})
		Convey("can Round Double", func() {
			example := types.NewDouble(4.006)
			res := example.Round(types.NewInteger(0))
			So(res.AsFloat64(), ShouldEqual, 4)
			example = types.NewDouble(4.006)
			res = example.Round(types.NewInteger(2))
			So(res.AsFloat64(), ShouldEqual, 4.01)
		})
		Convey("can create Random Double", func() {
			min := types.NewDouble(1)
			max := types.NewDouble(10)
			z := types.RandomDouble(min, max)
			So(z.AsFloat64(), ShouldBeBetween, 0, 11)
		})
	})
}
