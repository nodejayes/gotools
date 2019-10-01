package types

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewDateTime(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	g.Expect(example.Year().AsInt()).
		To(gomega.Equal(2019))
	g.Expect(example.Month().AsInt()).
		To(gomega.Equal(5))
	g.Expect(example.Day().AsInt()).
		To(gomega.Equal(1))
	g.Expect(example.Hour().AsInt()).
		To(gomega.Equal(22))
	g.Expect(example.Minute().AsInt()).
		To(gomega.Equal(45))
	g.Expect(example.Second().AsInt()).
		To(gomega.Equal(10))
	g.Expect(example.Millisecond().AsInt()).
		To(gomega.Equal(0))
}

func TestDateTime_SetYear(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetYear(*NewNumber(2))
	g.Expect(example.Year().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetMonth(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetMonth(*NewNumber(2))
	g.Expect(example.Month().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetDay(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetDay(*NewNumber(2))
	g.Expect(example.Day().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetHour(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetHour(*NewNumber(2))
	g.Expect(example.Hour().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetMinute(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetMinute(*NewNumber(2))
	g.Expect(example.Minute().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetSecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetSecond(*NewNumber(2))
	g.Expect(example.Second().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_SetMillisecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(22),
		*NewNumber(45),
		*NewNumber(10),
		ZERO)
	example.SetMillisecond(*NewNumber(2))
	g.Expect(example.Millisecond().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_AddYears(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddYears(*NewNumber(5))
	g.Expect(base.Year().AsInt()).To(gomega.Equal(2024))
}

func TestDateTime_AddMonths(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddMonths(*NewNumber(1))
	g.Expect(base.Month().AsInt()).To(gomega.Equal(6))
	base.AddMonths(*NewNumber(-1))
	base.AddMonths(*NewNumber(11))
	g.Expect(base.Month().AsInt()).To(gomega.Equal(4))
	base.AddMonths(*NewNumber(-11))
	base.AddMonths(*NewNumber(11))
	g.Expect(base.Year().AsInt()).To(gomega.Equal(2020))
}

func TestDateTime_AddDays(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	noLeapYear := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(2),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	leapYear := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2020),
		*NewNumber(2),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)

	base.AddDays(*NewNumber(1))
	g.Expect(base.Day().AsInt()).To(gomega.Equal(2))
	base.AddDays(*NewNumber(-1))
	base.AddDays(*NewNumber(31))
	g.Expect(base.Day().AsInt()).To(gomega.Equal(1))
	g.Expect(base.Month().AsInt()).To(gomega.Equal(6))
	base.AddDays(*NewNumber(-31))
	base.AddDays(*NewNumber(61))
	g.Expect(base.Day().AsInt()).To(gomega.Equal(1))
	g.Expect(base.Month().AsInt()).To(gomega.Equal(7))

	noLeapYear.AddDays(*NewNumber(27))
	g.Expect(noLeapYear.Month().AsInt()).To(gomega.Equal(2))
	g.Expect(noLeapYear.Day().AsInt()).To(gomega.Equal(28))
	noLeapYear.AddDays(*NewNumber(-27))
	noLeapYear.AddDays(*NewNumber(28))
	g.Expect(noLeapYear.Month().AsInt()).To(gomega.Equal(3))
	g.Expect(noLeapYear.Day().AsInt()).To(gomega.Equal(1))

	leapYear.AddDays(*NewNumber(28))
	g.Expect(leapYear.Month().AsInt()).To(gomega.Equal(2))
	g.Expect(leapYear.Day().AsInt()).To(gomega.Equal(29))
	leapYear.AddDays(*NewNumber(-28))
	leapYear.AddDays(*NewNumber(29))
	g.Expect(leapYear.Month().AsInt()).To(gomega.Equal(3))
	g.Expect(leapYear.Day().AsInt()).To(gomega.Equal(1))
}

func TestDateTime_AddHours(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddHours(*NewNumber(2))
	g.Expect(base.Hour().AsInt()).To(gomega.Equal(17))
	base.AddHours(*NewNumber(-2))
	base.AddHours(*NewNumber(24))
	g.Expect(base.Hour().AsInt()).To(gomega.Equal(15))
	g.Expect(base.Day().AsInt()).To(gomega.Equal(2))
}

func TestDateTime_AddMinutes(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddMinutes(*NewNumber(2))
	g.Expect(base.Minute().AsInt()).To(gomega.Equal(20))
	base.AddMinutes(*NewNumber(-2))
	base.AddMinutes(*NewNumber(62))
	g.Expect(base.Minute().AsInt()).To(gomega.Equal(20))
	g.Expect(base.Hour().AsInt()).To(gomega.Equal(16))
}

func TestDateTime_AddSeconds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddSeconds(*NewNumber(1))
	g.Expect(base.Second().AsInt()).To(gomega.Equal(3))
	base.AddSeconds(*NewNumber(-1))
	base.AddSeconds(*NewNumber(62))
	g.Expect(base.Second().AsInt()).To(gomega.Equal(4))
	g.Expect(base.Minute().AsInt()).To(gomega.Equal(19))
}

func TestDateTime_AddMilliseconds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	base := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(15),
		*NewNumber(18),
		*NewNumber(2),
		ZERO)
	base.AddMilliseconds(*NewNumber(200))
	g.Expect(base.Millisecond().AsInt()).To(gomega.Equal(200))
	base.AddMilliseconds(*NewNumber(-200))
	base.AddMilliseconds(*NewNumber(1205))
	g.Expect(base.Millisecond().AsInt()).To(gomega.Equal(205))
	g.Expect(base.Second().AsInt()).To(gomega.Equal(3))
}

func TestEmptyDateTime(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(EmptyDateTime().IsValid()).To(gomega.BeFalse())
}

func TestNewDateTimeInvalidLocation(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewDateTime(
		*NewString("notexists"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(18),
		*NewNumber(15),
		*NewNumber(1),
		*NewNumber(1),
	).IsValid()).To(gomega.BeFalse())
}

func TestDateTime_Clone(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	original := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(18),
		*NewNumber(15),
		*NewNumber(1),
		*NewNumber(1),
	)
	clone := original.Clone()
	g.Expect(clone == original).To(gomega.BeFalse())
}

func TestDateTime_ToZone(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	UTCDate := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(18),
		*NewNumber(15),
		*NewNumber(1),
		*NewNumber(1),
	)
	CETDate := UTCDate.ToZone(*NewString("CET"))
	g.Expect(CETDate.Hour().AsInt()).To(gomega.Equal(20))

	UTCtmp := CETDate.ToZone(*NewString("UTC"))
	g.Expect(UTCtmp.Hour().AsInt()).To(gomega.Equal(18))

	UTCtmp2 := CETDate.ToZone(*NewString("notexists"))
	g.Expect(UTCtmp2.IsValid()).To(gomega.BeFalse())
}

func TestDateTime_Equals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
		*NewString("UTC"),
		*NewNumber(2019),
		*NewNumber(5),
		*NewNumber(1),
		*NewNumber(18),
		*NewNumber(15),
		*NewNumber(1),
		*NewNumber(1),
	)
	example2 := example.Clone()
	example2.AddSeconds(*NewNumber(10))
	g.Expect(example.Equals(*example)).To(gomega.BeTrue())
	g.Expect(example.Equals(*example2)).To(gomega.BeFalse())
}
