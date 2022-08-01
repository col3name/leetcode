package _172

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(3))
	assert.Equal(t, 1, trailingZeroes(5))
	assert.Equal(t, 0, trailingZeroes(0))
	assert.Equal(t, 4, trailingZeroes(15))
	assert.Equal(t, 4, trailingZeroes(20))
	assert.Equal(t, 6, trailingZeroes(25))
}
