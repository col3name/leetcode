package _139

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))
	assert.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	assert.Equal(t, true, wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
	assert.Equal(t, true, wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	assert.Equal(t, false, wordBreak("aaaaaaa", []string{"aaaa", "aa"}))
	assert.Equal(t, false, wordBreak("a", []string{"b"}))
}

func Test36(t *testing.T) {
	assert.Equal(t, false, wordBreak("a", []string{"b"}))
}
