package sequencing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistance(t *testing.T) {
	d := LevenshteinDistance([]byte("kitten"), []byte("sitting"), nil, nil)
	assert.Equal(t, 3, d)
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	b.Run("6 char string", func(b *testing.B) {
		x := []byte("ABCDEF")
		y := []byte("ABCCDEF")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("25-30 char string", func(b *testing.B) {
		x := []byte("This is a longer string")
		y := []byte("This is a much  longer string")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("Long log line", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")

		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, nil, nil)
		}
	})
	b.Run("Long log line buffer reuse", func(b *testing.B) {
		x := []byte("10__8__0__146 kernel process Google Chrome Ca[3955] caught causing excessive wakeups. Observed wakeups rate (per sec): 392; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 317314")
		y := []byte("10__8__0__146 kernel process Sublime Text[802] caught causing excessive wakeups. Observed wakeups rate (per sec): 233; Maximum permitted wakeups rate (per sec): 150; Observation period: 300 seconds; Task lifetime number of wakeups: 95333")
		n := len(y)
		v0 := make([]int, n+1, n+1)
		v1 := make([]int, n+1, n+1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			LevenshteinDistance(x, y, v0, v1)
		}
	})

}
