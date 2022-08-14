package _1334

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	assert.Equal(t, 3, findTheCity(4, edges, 4))
}

func TestName2(t *testing.T) {
	edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	assert.Equal(t, 0, findTheCity(5, edges, 2))
}
