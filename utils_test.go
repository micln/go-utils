package go_utils

import (
	"testing"

	. "github.com/micln/go-utils/test"
)

func TestCamelCase(t *testing.T) {
	AssertEqual(t, CamelCase(`a_b_c`), `ABC`)
}

func TestUnderlineCase(t *testing.T) {
	AssertEqual(t, UnderlineCase(`ABC`), `a_b_c`)
}
