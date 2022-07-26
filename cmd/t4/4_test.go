package t4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 2, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
