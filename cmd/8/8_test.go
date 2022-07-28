package t8

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("   -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 3, myAtoi("3.14159"))
	assert.Equal(t, 13, myAtoi("13.14159"))
	assert.Equal(t, 0, myAtoi("0.14159"))
	assert.Equal(t, 0, myAtoi("0.5"))
	assert.Equal(t, 0, myAtoi("+-12"))
	assert.Equal(t, 2147483647, myAtoi("21474836460"))
	assert.Equal(t, -12, myAtoi("  -0012a42"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
	assert.Equal(t, 12345678, myAtoi("  0000000000012345678"))
	assert.Equal(t, 0, myAtoi("00000-42a1234"))
	atoi, err := strconv.Atoi("0.5")
	fmt.Println(err)
	assert.Equal(t, 0, atoi)
}

func TestName2(t *testing.T) {
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
}
