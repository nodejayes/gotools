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

func TestString_ToUpperCase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").ToUpperCase().AsString()).
		To(gomega.Equal("ABC"))
	g.Expect(NewString("Abc").ToUpperCase().AsString()).
		To(gomega.Equal("ABC"))
	g.Expect(NewString("ABC").ToUpperCase().AsString()).
		To(gomega.Equal("ABC"))
}

func TestString_ToLowerCase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("ABC").ToLowerCase().AsString()).
		To(gomega.Equal("abc"))
	g.Expect(NewString("aBc").ToLowerCase().AsString()).
		To(gomega.Equal("abc"))
	g.Expect(NewString("abc").ToLowerCase().AsString()).
		To(gomega.Equal("abc"))
}

func TestString_Equal(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").Equal(*NewString("abc"))).
		To(gomega.BeTrue())
	g.Expect(NewString("abcd").Equal(*NewString("abc"))).
		To(gomega.BeFalse())
}

func TestString_Pad(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	l := *NewNumber(8)
	g.Expect(NewString("abc").Pad(l, *NewString("_-")).AsString()).
		To(gomega.Equal("_-abc_-_"))
	g.Expect(NewString("abcdefgh").Pad(l, *NewString("_-")).AsString()).
		To(gomega.Equal("abcdefgh"))
}

func TestString_PadLeft(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	l := *NewNumber(8)
	g.Expect(NewString("abc").PadLeft(l, *NewString(" ")).AsString()).
		To(gomega.Equal("     abc"))
	g.Expect(NewString("abcabcab").PadLeft(l, *NewString(" ")).AsString()).
		To(gomega.Equal("abcabcab"))
	g.Expect(NewString("abcabcabc").PadLeft(l, *NewString(" ")).AsString()).
		To(gomega.Equal("abcabcabc"))
}

func TestString_PadRight(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	l := *NewNumber(8)
	g.Expect(NewString("abc").PadRight(l, *NewString(" ")).AsString()).
		To(gomega.Equal("abc     "))
	g.Expect(NewString("abcabcab").PadRight(l, *NewString(" ")).AsString()).
		To(gomega.Equal("abcabcab"))
	g.Expect(NewString("abcabcabc").PadRight(l, *NewString(" ")).AsString()).
		To(gomega.Equal("abcabcabc"))
}

func TestString_Concat(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("a").Concat(*NewString("b")).AsString()).
		To(gomega.Equal("ab"))
}

func TestString_Repeat(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	times := NewNumber(2)
	g.Expect(NewString("abc").Repeat(*times).AsString()).
		To(gomega.Equal("abcabc"))
}

func TestString_Split(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	splitter := *NewString("b")
	splitted := NewString("abcabc").Split(splitter)
	g.Expect(splitted).To(gomega.HaveLen(3))
}

func TestString_TextBetween(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	tmp := NewString("abcdefg").
		TextBetween(*NewString("a"), *NewString("c"))
	tmp2 := NewString("abcdefgabcdefg").
		TextBetween(*NewString("a"), *NewString("c"))
	tmp3 := NewString("abbcbbcdefgabcdefg").
		TextBetween(*NewString("a"), *NewString("c"))
	g.Expect(len(tmp)).
		To(gomega.Equal(1))
	g.Expect(tmp[0].AsString()).To(gomega.Equal("b"))

	g.Expect(len(tmp2)).
		To(gomega.Equal(2))
	g.Expect(tmp2[0].AsString()).To(gomega.Equal("b"))
	g.Expect(tmp2[1].AsString()).To(gomega.Equal("b"))

	g.Expect(len(tmp3)).
		To(gomega.Equal(2))
	g.Expect(tmp3[0].AsString()).To(gomega.Equal("bb"))
	g.Expect(tmp3[1].AsString()).To(gomega.Equal("b"))
}

func TestString_IndexOf(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").IndexOf(*NewString("b")).AsInt()).
		To(gomega.Equal(1))
	g.Expect(NewString("abc").IndexOf(*NewString("x")).AsInt()).
		To(gomega.Equal(-1))
	g.Expect(NewString("abcabcabcabc").IndexOf(*NewString("b")).AsInt()).
		To(gomega.Equal(1))
}

func TestString_LastIndexOf(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").LastIndexOf(*NewString("b")).AsInt()).
		To(gomega.Equal(1))
	g.Expect(NewString("abc").LastIndexOf(*NewString("x")).AsInt()).
		To(gomega.Equal(-1))
	g.Expect(NewString("abcabcabcabc").LastIndexOf(*NewString("b")).AsInt()).
		To(gomega.Equal(10))
}

func TestString_SubString(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abcdefgh").SubString(*NewNumber(1), *NewNumber(2)).AsString()).
		To(gomega.Equal("bc"))
	g.Expect(NewString("abcdefgh").SubString(*NewNumber(1), *NewNumber(-1)).AsString()).
		To(gomega.Equal("bcdefgh"))
}

func TestString_Remove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abbbbc").Remove(*NewNumber(1), *NewNumber(3)).AsString()).
		To(gomega.Equal("abc"))
	g.Expect(NewString("abbbbc").Remove(*NewNumber(1), *NewNumber(20)).AsString()).
		To(gomega.Equal("a"))
	g.Expect(NewString("abbbbc").Remove(*NewNumber(-1), *NewNumber(20)).AsString()).
		To(gomega.Equal(""))
}

func TestString_Insert(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(NewString("abc").Insert(*NewNumber(1), *NewString("x")).AsString()).
		To(gomega.Equal("axbc"))
	g.Expect(NewString("abc").Insert(*NewNumber(0), *NewString("x")).AsString()).
		To(gomega.Equal("xabc"))
	g.Expect(NewString("abc").Insert(*NewNumber(3), *NewString("x")).AsString()).
		To(gomega.Equal("abcx"))
	g.Expect(NewString("abc").Insert(*NewNumber(-1), *NewString("x")).AsString()).
		To(gomega.Equal("xabc"))
	g.Expect(NewString("abc").Insert(*NewNumber(4), *NewString("x")).AsString()).
		To(gomega.Equal("abcx"))
}
