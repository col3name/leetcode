package _150

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
}
