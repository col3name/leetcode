package t7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
}
