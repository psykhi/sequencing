package sequencing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHirschberg(t *testing.T) {
	a := []byte("ABCDEF")
	b := []byte("ABCCDEF")
	//
	//a := []byte("This is a longer string")
	//b := []byte("This is a much longer string")

	z, w := Hirschberg(a, b)
	fmt.Printf("%s\n%s\n", string(z), string(w))
	assert.Equal(t, "ABC-DEF", string(z))
	assert.Equal(t, "ABCCDEF", string(w))
}

func BenchmarkHirschberg(b *testing.B) {
	b.Run("6 char string", func(b *testing.B) {
		x := []byte("ABCDEF")
		y := []byte("ABCCDEF")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Hirschberg(x, y)
		}
	})
	b.Run("25-30 char string", func(b *testing.B) {
		x := []byte("This is a longer string")
		y := []byte("This is a much longer string")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Hirschberg(x, y)
		}
	})
	b.Run("Long log line", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Hirschberg(x, y)
		}
	})
}
