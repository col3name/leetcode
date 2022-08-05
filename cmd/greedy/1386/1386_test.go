package _1386

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 4, maxNumberOfFamilies(3, [][]int{
		{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10},
	}))
	assert.Equal(t, 2, maxNumberOfFamilies(2, [][]int{
		{2, 1}, {1, 8}, {2, 6},
	}))
	assert.Equal(t, 4, maxNumberOfFamilies(4, [][]int{
		{4, 3}, {1, 4}, {4, 6}, {1, 7},
	}))
	assert.Equal(t, 5, maxNumberOfFamilies(3, [][]int{{2, 3}}))
}
