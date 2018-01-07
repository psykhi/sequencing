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
