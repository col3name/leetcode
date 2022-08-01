package _380

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	set := Constructor()
	assert.Equal(t, true, set.Insert(1))
	assert.Equal(t, false, set.Remove(2))
	assert.Equal(t, true, set.Insert(2))
	random := set.GetRandom()
	assert.True(t, random == 1 || random == 2)
	assert.Equal(t, true, set.Remove(1))
	assert.Equal(t, false, set.Insert(2))
	random = set.GetRandom()
	assert.Equal(t, 2, random)
}
