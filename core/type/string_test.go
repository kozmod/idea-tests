package _type

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTrimString(t *testing.T) {
	str1 := "!! X !!"
	str2 := "@@Y$$"
	str3 := "   Z "

	// Trimming the given strings
	// Using Trim() function
	res1 := strings.Trim(str1, "!")
	res2 := strings.Trim(str2, "@$")
	res3 := strings.Trim(str3, " ")

	assert.Equal(t, " X ", res1)
	assert.Equal(t, "Y", res2)
	assert.Equal(t, "Z", res3)
}
