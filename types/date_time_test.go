package types

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewDateTime(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	example := NewDateTime(
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
