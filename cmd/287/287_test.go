package _87

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
	assert.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
}
