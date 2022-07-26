package _29

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 4, longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))
}
