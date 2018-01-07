package alignment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeedlemanWunsch(t *testing.T) {
	a := []byte("ABCDEF")
	b := []byte("ABCCDEF")

	z, w := NeedlemanWunsch(a, b)
	fmt.Printf("%s\n%s\n", string(z), string(w))
	assert.Equal(t, "AB-CDEF", string(z))
	assert.Equal(t, "ABCCDEF", string(w))
}

func BenchmarkNeedlemanWunsch(b *testing.B) {
	b.Run("6 char string", func(b *testing.B) {
		x := []byte("ABCDEF")
		y := []byte("ABCCDEF")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NeedlemanWunsch(x, y)
		}
	})
	b.Run("25-30 char string", func(b *testing.B) {
		x := []byte("This is a longer string")
		y := []byte("This is a much  longer string")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NeedlemanWunsch(x, y)
		}
	})
}
