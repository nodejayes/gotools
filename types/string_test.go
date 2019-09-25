package types

import (
	"github.com/onsi/gomega"
	"testing"
	"time"
)

func TestNewStringFromInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(1).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(1).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromInt16(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(int16(1)).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(int16(1)).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromInt32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(int32(1)).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(int32(1)).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromInt64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(int64(1)).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(int64(1)).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromFloat32(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(float32(1.5)).AsString()).
		To(gomega.Equal("1.5"))
	g.Expect(NewString(float32(1)).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(float32(1.5)).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromFloat64(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(1.5).AsString()).
		To(gomega.Equal("1.5"))
	g.Expect(NewString(float64(1)).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString(1.5).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("1").AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString("1").IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromByteArray(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString([]byte("1")).AsString()).
		To(gomega.Equal("1"))
	g.Expect(NewString([]byte("1")).IsValid()).
		To(gomega.BeTrue())
}

func TestNewStringFromNumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(*NewNumber(1)).AsString()).
		To(gomega.Equal("1"))
}

func TestNewStringFromNumberPointer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(NewNumber(1)).AsString()).
		To(gomega.Equal("1"))
}

func TestNewString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(*NewString("1")).AsString()).
		To(gomega.Equal("1"))
}

func TestNewStringFromStringPointer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(NewString("1")).AsString()).
		To(gomega.Equal("1"))
}

func TestNewStringFromCharUint8(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(NewString("1"[0])).AsString()).
		To(gomega.Equal("1"))
}

func TestNewStringNoPanicAndDefaultWhenNotSupported(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString(time.Now()).AsString()).
		To(gomega.Equal(""))
	g.Expect(NewString(time.Now()).IsValid()).
		To(gomega.BeFalse())
}

func TestString_AsString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").AsString()).
		To(gomega.Equal("abc"))
}

func TestString_AsNumber(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("1").AsNumber().AsInt()).
		To(gomega.Equal(1))
}

func TestString_IsValid(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("1").IsValid()).
		To(gomega.BeTrue())
	g.Expect(NewString(time.Now()).IsValid()).
		To(gomega.BeFalse())
}

func TestString_Clone(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	original := NewString("1")
	clone := original.Clone()
	g.Expect(original == clone).
		To(gomega.BeFalse())
}

func TestString_Length(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("12345").Length().AsInt()).
		To(gomega.Equal(5))
}

func TestString_CharAt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").CharAt(*NewNumber(1)).AsString()).
		To(gomega.Equal("a"))
	g.Expect(NewString("abc").CharAt(*NewNumber(3)).AsString()).
		To(gomega.Equal("c"))
	g.Expect(NewString("abc").CharAt(*NewNumber(4)).AsString()).
		To(gomega.Equal(""))
	g.Expect(NewString("abc").CharAt(*NewNumber(4)).IsValid()).
		To(gomega.BeFalse())
	g.Expect(NewString("abc").CharAt(*NewNumber(0)).AsString()).
		To(gomega.Equal(""))
	g.Expect(NewString("abc").CharAt(*NewNumber(0)).IsValid()).
		To(gomega.BeFalse())
}
