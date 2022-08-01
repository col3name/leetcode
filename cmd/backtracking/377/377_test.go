package _377

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 7, combinationSum4([]int{1, 2, 3}, 4))
	//assert.Equal(t, 0, combinationSum4([]int{9}, 3))
}
