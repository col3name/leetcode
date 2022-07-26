package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"
	assert.Equal(t, true, exist(board, word))
}
