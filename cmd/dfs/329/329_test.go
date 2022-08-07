package _329

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	assert.Equal(t, 4, longestIncreasingPath(matrix))
	assert.Equal(t, 2, longestIncreasingPath([][]int{{1, 2}}))
}
