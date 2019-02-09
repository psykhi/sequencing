package sequencing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCS(t *testing.T) {
	d := LCS([]byte("kitten"), []byte("sitting"), nil, nil)
	assert.Equal(t, 4, d)

	d = LCS([]byte("MZJAWXU"), []byte("XMJYAUZ"), nil, nil)
	assert.Equal(t, 4, d)

	d = LCS([]byte("AGCAT"), []byte("GAC"), nil, nil)
	assert.Equal(t, 2, d)

	d = LCSStrings([]string{"hello", "world"}, []string{"hello", "universe", "something"}, nil, nil)
	assert.Equal(t, 1, d)
}
