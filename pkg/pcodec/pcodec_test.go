package pcodec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	a := []uint64{
		65465467874598,
		56496876965878,
		48974654125465,
		87897441897545,
	}
	str, err := Encode(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func BenchmarkEncode(b *testing.B) {
	a := []uint64{
		65465467874598,
		56496876965878,
		48974654125465,
		87897441897545,
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Encode(a)
	}
}

func TestDecode(t *testing.T) {
	a := []uint64{
		65465467874598,
		56496876965878,
		48974654125465,
		87897441897545,
	}
	str, err := Encode(a)
	if err != nil {
		t.Fatal(err)
	}

	b, der := Decode(str)
	if der != nil {
		t.Fatal(der)
	}
	assert.Equal(t, a, b)
	t.Log(a[0] & 0x4)
	t.Log(a[0] & 0x8)
}

func BenchmarkDecode(b *testing.B) {
	a := []uint64{
		65465467874598,
		56496876965878,
		48974654125465,
		87897441897545,
	}
	str, err := Encode(a)
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(str)
	}
}
