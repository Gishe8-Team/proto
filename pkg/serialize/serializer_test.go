package serialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	a := []uint64{
		65465467874598,
		56496876965878,
		48974654125465,
		87897441897545,
	}
	data, err := Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x", data)
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	a := []uint64{
		65465467,
		56496876,
		48974654,
		87897441,
	}
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(a)
	}
}

func TestUnmarshal(t *testing.T) {
	a := []uint64{
		65465467,
		56496876,
		48974654,
		87897441,
	}
	data, err := Marshal(a)
	if err != nil {
		t.Fatal(err)
	}

	b, ber := Unmarshal(data)
	if ber != nil {
		t.Fatal(err)
	}

	assert.Equal(t, a, b)
	t.Logf("data: %#v", b)
}

func BenchmarkUnmarshal(b *testing.B) {
	a := []uint64{
		65465467,
		56496876,
		48974654,
		87897441,
	}
	data, _ := Marshal(a)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = Unmarshal(data)
	}
}
