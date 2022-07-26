package _07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.True(t, canFinish(2, [][]int{{0, 1}}))
	assert.True(t, canFinish(2, [][]int{{1, 0}}))
	assert.False(t, canFinish(2, [][]int{{0, 1}, {1, 0}}))
}
