package _735

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{5, 10}, asteroidCollision([]int{5, 10, -5})))
	assert.True(t, reflect.DeepEqual([]int{}, asteroidCollision([]int{8, -8})))
	assert.True(t, reflect.DeepEqual([]int{10}, asteroidCollision([]int{10, 2 - 5})))
}
