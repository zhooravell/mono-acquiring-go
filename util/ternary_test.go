// revive:disable:var-naming
package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	assert.Equal(t, "first", Ternary(true, "first", "second"))
	assert.Equal(t, "second", Ternary(false, "first", "second"))

	assert.Equal(t, 1, Ternary(true, 1, 2))
	assert.Equal(t, 2, Ternary(false, 1, 2))

	assert.Equal(t, 1.1, Ternary(true, 1.1, 2.2))
	assert.Equal(t, 2.2, Ternary(false, 1.1, 2.2))
}
