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
	g.Expect(example.SetYear(*NewNumber(2)).Year().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetMonth(*NewNumber(2)).Month().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetDay(*NewNumber(2)).Day().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetHour(*NewNumber(2)).Hour().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetMinute(*NewNumber(2)).Minute().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetSecond(*NewNumber(2)).Second().AsInt()).To(gomega.Equal(2))
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
	g.Expect(example.SetMillisecond(*NewNumber(2)).Millisecond().AsInt()).To(gomega.Equal(2))
}
