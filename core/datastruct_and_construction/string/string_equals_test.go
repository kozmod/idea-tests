package string

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEqualsIgnoreCase(t *testing.T) {
	assert.True(t, strings.EqualFold("abc", "ABC"))
	assert.True(t, strings.EqualFold("abc", "aBC"))
	assert.True(t, strings.EqualFold("abc", "AbC"))
	assert.False(t, strings.EqualFold("abc", "AbCd"))
}
