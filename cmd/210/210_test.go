package _10

import (
	"github.com/col3name/algo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	result := findOrder(2, [][]int{{0, 1}})
	assert.True(t, util.Equal(result, []int{0, 1}))
}
