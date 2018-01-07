package alignment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHirschberg(t *testing.T) {
	a := []byte("ABCDEF")
	b := []byte("ABCCDEF")

	z, w := Hirschberg(a, b)
	fmt.Printf("%s\n%s\n", string(z), string(w))
	assert.Equal(t, "AB-CDEF", string(z))
	assert.Equal(t, "ABCCDEF", string(w))
}
