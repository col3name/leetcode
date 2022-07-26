package _20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}))
}
