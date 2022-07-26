package _004

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 6, longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
