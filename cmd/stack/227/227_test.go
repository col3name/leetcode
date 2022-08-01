package _227

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 7, calculate("3+2*2"))
	assert.Equal(t, 1, calculate(" 3/2 "))
	assert.Equal(t, 5, calculate(" 3+5 / 2 "))
}
