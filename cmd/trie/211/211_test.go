package _11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	assert.Equal(t, false, obj.Search("pad"))
	assert.Equal(t, true, obj.Search("bad"))
	assert.Equal(t, true, obj.Search(".ad"))
	assert.Equal(t, true, obj.Search("b.."))
}
