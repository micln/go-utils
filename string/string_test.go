package string

import (
	"testing"

	"docker/pkg/testutil/assert"
)

func TestString_Equal(t *testing.T) {
	assert.Equal(t, New(`abc`).Equal([]byte{'a', 'b', 'c'}), true)
}

func TestString_Contains(t *testing.T) {
	assert.Equal(t, New(`abc`).Contains(`b`), true)
	assert.Equal(t, New(`abc`).Contains(`d`), false)

	assert.Equal(t, New([]byte{'a', 'b', 'c'}).Contains(`b`), true)
}

func TestString_Trim(t *testing.T) {
	assert.Equal(t, New(` abc `).Trim(` `).Equal(`abc`),true)
}