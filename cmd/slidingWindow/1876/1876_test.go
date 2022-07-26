package _876

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 4, countGoodSubstrings("aababcabc"))
}
