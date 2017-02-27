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
	assert.Equal(t, New([]rune{'李', '二', '狗'}).Contains(`二`), true)
}

func TestString_Trim(t *testing.T) {
	assert.Equal(t, New(` abc `).Trim(` `).Equal(`abc`), true)
	assert.Equal(t, New(`李二狗`).Trim(`狗`).Equal(`李二`), true)
}

func TestString_Replace(t *testing.T) {
	assert.Equal(t, New(`aabc`).Replace(`a`, `b`).Equal(`bbbc`), true)
	assert.Equal(t, New(`aabc`).Replace(`a`, `b`, 1).Equal(`babc`), true)
	assert.Equal(t, New(`李二狗`).Replace(`二`, `三`).Equal(`李三狗`), true)
}
