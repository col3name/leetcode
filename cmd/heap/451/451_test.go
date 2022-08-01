package _451

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, "eert", frequencySort("tree"))
	assert.Equal(t, "aaaccc", frequencySort("cccaaa"))
	assert.Equal(t, "bbAa", frequencySort("Aabb"))
	assert.Equal(t, "eeee", frequencySort("eeee"))
}
