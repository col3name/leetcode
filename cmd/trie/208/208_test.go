package _08

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	word := "hello"
	obj := Constructor()
	obj.Insert(word)
	fmt.Println(obj.root)
	assert.Equal(t, true, obj.Search(word))
	assert.Equal(t, true, obj.StartsWith("hell"))
	assert.Equal(t, false, obj.StartsWith("hellos"))
}

func TestName2(t *testing.T) {
	word := "a"
	obj := Constructor()
	obj.Insert(word)
	fmt.Println(obj.root)
	assert.Equal(t, true, obj.Search(word))
	assert.Equal(t, true, obj.StartsWith("a"))
}

func TestName4(t *testing.T) {
	pair := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for key, val := range pair {
		fmt.Println(key, val)
	}

}
