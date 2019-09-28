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
	example1, _ := NewTimeSpanFromISOString(*NewString("1.22:05:08 200"))
	example2, _ := NewTimeSpanFromISOString(*NewString("22:05:08 200"))
	example3, _ := NewTimeSpanFromISOString(*NewString("22:05:08"))
	example4, _ := NewTimeSpanFromISOString(*NewString("1.22:05:08"))
	example5, _ := NewTimeSpanFromISOString(*NewString("1."))
	example6, err := NewTimeSpanFromISOString(*NewString("-"))

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
