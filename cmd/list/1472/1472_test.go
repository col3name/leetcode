package _472

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type URL string

var (
	Google        URL = "https://google.com"
	Github        URL = "https://github.com"
	Stackoverflow URL = "https://stackoverflow.com"
	Medium        URL = "https://medium.com"
	Reddit        URL = "https://reddit.com"
	Leetcode      URL = "https://leetcode.com"
	Facebook      URL = "https://facebook.com"
	Youtube       URL = "https://youtube.com"
	Linkedin      URL = "https://linkedin.com"
)

func TestName(t *testing.T) {
	history := Constructor(string(Google))
	history.Visit(string(Github))
	history.Visit(string(Stackoverflow))
	history.Visit(string(Medium))
	history.Visit(string(Reddit))
	assert.Equal(t, string(Medium), history.Back(1))
	assert.Equal(t, string(Reddit), history.Forward(1))
	assert.Equal(t, string(Google), history.Back(6))
	assert.Equal(t, string(Reddit), history.Forward(history.Size+1))
}

func TestName2(t *testing.T) {
	history := Constructor(string(Google))
	history.Visit(string(Github))
	history.Visit(string(Stackoverflow))
	history.Visit(string(Medium))

	assert.Equal(t, string(Stackoverflow), history.Back(1))
	assert.Equal(t, string(Github), history.Back(1))
	assert.Equal(t, string(Stackoverflow), history.Forward(1))
	assert.Equal(t, string(Medium), history.Forward(2))
	history.Visit(string(Reddit))
	assert.Equal(t, string(Reddit), history.Forward(2))
	assert.Equal(t, string(Stackoverflow), history.Back(2))
	assert.Equal(t, string(Google), history.Back(7))
}

func TestName3(t *testing.T) {
	history := Constructor(string(Leetcode))
	history.Visit(string(Google))
	history.Visit(string(Facebook))
	history.Visit(string(Youtube))

	assert.Equal(t, string(Facebook), history.Back(1))
	assert.Equal(t, string(Google), history.Back(1))
	assert.Equal(t, string(Facebook), history.Forward(1))

	history.Visit(string(Linkedin))
	assert.Equal(t, string(Linkedin), history.Forward(2))
	assert.Equal(t, string(Google), history.Back(2))
	assert.Equal(t, string(Leetcode), history.Back(7))
}

func TestName4(t *testing.T) {
	history := Constructor(string(Leetcode))
	history.Visit(string(Google))
	history.Visit(string(Facebook))
	history.Visit(string(Youtube))

	assert.Equal(t, string(Facebook), history.Back(1))
	assert.Equal(t, string(Google), history.Back(1))
	assert.Equal(t, string(Facebook), history.Forward(1))

	history.Visit(string(Linkedin))
	assert.Equal(t, string(Linkedin), history.Forward(2))
	assert.Equal(t, string(Google), history.Back(2))
	assert.Equal(t, string(Leetcode), history.Back(7))
}
