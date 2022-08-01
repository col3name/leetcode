package _032

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	obj := Constructor([]string{"cd", "f", "kl"})
	assert.Equal(t, false, obj.Query('a'))
	assert.Equal(t, false, obj.Query('b'))
	assert.Equal(t, false, obj.Query('c'))
	assert.Equal(t, true, obj.Query('d'))
	assert.Equal(t, false, obj.Query('e'))
	assert.Equal(t, true, obj.Query('f'))
	assert.Equal(t, false, obj.Query('g'))
	assert.Equal(t, false, obj.Query('h'))
	assert.Equal(t, false, obj.Query('i'))
	assert.Equal(t, false, obj.Query('j'))
	assert.Equal(t, false, obj.Query('k'))
	assert.Equal(t, true, obj.Query('l'))
}
