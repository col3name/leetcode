package _1

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9)))
	assert.True(t, reflect.DeepEqual([]int{1, 2}, twoSum([]int{3, 2, 4}, 6)))
	assert.True(t, reflect.DeepEqual([]int{0, 1}, twoSum([]int{3, 3}, 6)))
	assert.True(t, reflect.DeepEqual([]int{}, twoSum([]int{3, 3}, 7)))
}
