package _3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "1", input: "LVIII", expected: 58},
		{name: "2", input: "III", expected: 3},
		{name: "3", input: "MDXI", expected: 1511},
		{name: "4", input: "MCMXCIV", expected: 1994},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, romanToInt(test.input))
			assert.Equal(t, test.expected, romanToInt2(test.input))
		})
	}
}
