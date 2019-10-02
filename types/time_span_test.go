package types

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewTimeSpan(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(
		*NewNumber(0),
		*NewNumber(1),
		*NewNumber(2),
		*NewNumber(5),
		*NewNumber(50),
	)

	g.Expect(example.IsValid()).To(gomega.BeTrue())
	g.Expect(example.Millisecond().AsInt()).To(gomega.Equal(50))
	g.Expect(example.Second().AsInt()).To(gomega.Equal(5))
	g.Expect(example.Minute().AsInt()).To(gomega.Equal(2))
	g.Expect(example.Hour().AsInt()).To(gomega.Equal(1))
	g.Expect(example.Day().AsInt()).To(gomega.Equal(0))
}

func TestNewTimeSpanFromMilliseconds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1 := NewTimeSpanFromMilliseconds(*NewNumber(1000))
	example2 := NewTimeSpanFromMilliseconds(*NewNumber(3665001))

	g.Expect(example1.IsValid()).To(gomega.BeTrue())
	g.Expect(example1.Millisecond().AsInt()).To(gomega.Equal(0))
	g.Expect(example1.Second().AsInt()).To(gomega.Equal(1))
	g.Expect(example1.Minute().AsInt()).To(gomega.Equal(0))
	g.Expect(example1.Hour().AsInt()).To(gomega.Equal(0))
	g.Expect(example1.Day().AsInt()).To(gomega.Equal(0))

	g.Expect(example2.IsValid()).To(gomega.BeTrue())
	g.Expect(example2.Millisecond().AsInt()).To(gomega.Equal(1))
	g.Expect(example2.Second().AsInt()).To(gomega.Equal(5))
	g.Expect(example2.Minute().AsInt()).To(gomega.Equal(1))
	g.Expect(example2.Hour().AsInt()).To(gomega.Equal(1))
	g.Expect(example2.Day().AsInt()).To(gomega.Equal(0))
}

func TestNewTimeSpanFromISOString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1, _ := NewTimeSpanFromString(*NewString("1.22:05:08 200"))
	example2, _ := NewTimeSpanFromString(*NewString("22:05:08 200"))
	example3, _ := NewTimeSpanFromString(*NewString("22:05:08"))
	example4, _ := NewTimeSpanFromString(*NewString("1.22:05:08"))
	example5, _ := NewTimeSpanFromString(*NewString("1."))
	example6, err := NewTimeSpanFromString(*NewString("-"))

	g.Expect(example1.IsValid()).To(gomega.BeTrue())
	g.Expect(example1.Millisecond().AsInt()).To(gomega.Equal(200))
	g.Expect(example1.Second().AsInt()).To(gomega.Equal(8))
	g.Expect(example1.Minute().AsInt()).To(gomega.Equal(5))
	g.Expect(example1.Hour().AsInt()).To(gomega.Equal(22))
	g.Expect(example1.Day().AsInt()).To(gomega.Equal(1))

	g.Expect(example2.IsValid()).To(gomega.BeTrue())
	g.Expect(example2.Millisecond().AsInt()).To(gomega.Equal(200))
	g.Expect(example2.Second().AsInt()).To(gomega.Equal(8))
	g.Expect(example2.Minute().AsInt()).To(gomega.Equal(5))
	g.Expect(example2.Hour().AsInt()).To(gomega.Equal(22))
	g.Expect(example2.Day().AsInt()).To(gomega.Equal(0))

	g.Expect(example3.IsValid()).To(gomega.BeTrue())
	g.Expect(example3.Millisecond().AsInt()).To(gomega.Equal(0))
	g.Expect(example3.Second().AsInt()).To(gomega.Equal(8))
	g.Expect(example3.Minute().AsInt()).To(gomega.Equal(5))
	g.Expect(example3.Hour().AsInt()).To(gomega.Equal(22))
	g.Expect(example3.Day().AsInt()).To(gomega.Equal(0))

	g.Expect(example4.IsValid()).To(gomega.BeTrue())
	g.Expect(example4.Millisecond().AsInt()).To(gomega.Equal(0))
	g.Expect(example4.Second().AsInt()).To(gomega.Equal(8))
	g.Expect(example4.Minute().AsInt()).To(gomega.Equal(5))
	g.Expect(example4.Hour().AsInt()).To(gomega.Equal(22))
	g.Expect(example4.Day().AsInt()).To(gomega.Equal(1))

	g.Expect(example5.IsValid()).To(gomega.BeTrue())
	g.Expect(example5.Millisecond().AsInt()).To(gomega.Equal(0))
	g.Expect(example5.Second().AsInt()).To(gomega.Equal(0))
	g.Expect(example5.Minute().AsInt()).To(gomega.Equal(0))
	g.Expect(example5.Hour().AsInt()).To(gomega.Equal(0))
	g.Expect(example5.Day().AsInt()).To(gomega.Equal(1))

	g.Expect(example6.IsValid()).To(gomega.BeFalse())
	g.Expect(err).To(gomega.Not(gomega.BeNil()))
}

func TestTimeSpan_AsString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1 := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	example2 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(0), *NewNumber(0), *NewNumber(0))
	g.Expect(example1.AsString().AsString()).To(gomega.Equal("365.22:45:05 200"))
	g.Expect(example2.AsString().AsString()).To(gomega.Equal("00:00:00"))
}

func TestTimeSpan_Day(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.Day().AsInt()).To(gomega.Equal(365))
}

func TestTimeSpan_Hour(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.Hour().AsInt()).To(gomega.Equal(22))
}

func TestTimeSpan_Minute(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.Minute().AsInt()).To(gomega.Equal(45))
}

func TestTimeSpan_Second(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.Second().AsInt()).To(gomega.Equal(5))
}

func TestTimeSpan_Millisecond(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.Millisecond().AsInt()).To(gomega.Equal(200))
}

func TestTimeSpan_IsValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.IsValid()).To(gomega.BeTrue())
	g.Expect(EmptyString().IsValid()).To(gomega.BeFalse())
}

func TestTimeSpan_TotalDays(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.TotalDays().AsFloat64()).To(gomega.Equal(365.9479768518519))
}

func TestTimeSpan_TotalHours(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(365), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.TotalHours().AsFloat64()).To(gomega.Equal(8782.751444444444))
}

func TestTimeSpan_TotalMinutes(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(22), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.TotalMinutes().AsFloat64()).To(gomega.Equal(1365.0866666666666))
}

func TestTimeSpan_TotalSeconds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(45), *NewNumber(5), *NewNumber(200))
	g.Expect(example.TotalSeconds().AsFloat64()).To(gomega.Equal(2705.2))
}

func TestTimeSpan_TotalMilliseconds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(0), *NewNumber(5), *NewNumber(200))
	g.Expect(example.TotalMilliseconds().AsInt()).To(gomega.Equal(5200))
}

func TestTimeSpan_Add(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(0), *NewNumber(0), *NewNumber(0))
	adder := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(5), *NewNumber(2), *NewNumber(0))
	g.Expect(example.Add(*adder).Minute().AsInt()).To(gomega.Equal(5))
	g.Expect(example.Add(*adder).Second().AsInt()).To(gomega.Equal(2))
}

func TestTimeSpan_Equals(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(0))
	example2 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(1))
	g.Expect(example1.Equals(*example2)).To(gomega.BeFalse())
	g.Expect(example2.Equals(*example2)).To(gomega.BeTrue())
}

func TestTimeSpan_IsAfter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(0))
	example2 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(1))
	g.Expect(example1.IsAfter(*example2)).To(gomega.BeFalse())
	g.Expect(example2.IsAfter(*example1)).To(gomega.BeTrue())
}

func TestTimeSpan_IsBefore(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example1 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(0))
	example2 := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(1))
	g.Expect(example1.IsBefore(*example2)).To(gomega.BeTrue())
	g.Expect(example2.IsBefore(*example1)).To(gomega.BeFalse())
}

func TestTimeSpan_Subtract(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(0))
	subtracted := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(5), *NewNumber(2), *NewNumber(0))
	g.Expect(example.Subtract(*subtracted).Minute().AsInt()).To(gomega.Equal(1))
	g.Expect(example.Subtract(*subtracted).Second().AsInt()).To(gomega.Equal(1))
}

func TestTimeSpan_Negate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewTimeSpan(*NewNumber(0), *NewNumber(0), *NewNumber(6), *NewNumber(3), *NewNumber(0))
	g.Expect(example.Negate().Minute().AsInt()).To(gomega.Equal(-6))
	g.Expect(example.Negate().Second().AsInt()).To(gomega.Equal(-3))
	g.Expect(example.Negate().Millisecond().AsInt()).To(gomega.Equal(0))
}
