package _91

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("226"))
	assert.Equal(t, 0, numDecodings("06"))
}
