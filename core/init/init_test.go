package init

import (
	"github.com/magiconair/properties/assert"
	"testing"
)
import _ "github.com/kozmod/idea-tests/core/init/sub"
import s "github.com/kozmod/idea-tests/core/init/sub"

func TestInit(t *testing.T) {
	root := getVal()
	sub := getVal()
	assert.Equal(t, Val, root)
	assert.Equal(t, s.Val, sub)
}
