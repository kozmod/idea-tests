package ref

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type testStruct struct {
	val string
}

func (s testStruct) noChangeVal(val string) {
	s.val = val
}

func (s *testStruct) changeVal(val string) {
	s.val = val
}

func noChangeTestStructVal(s testStruct, val string) {
	s.val = val
}

func changeTestStructVal(s *testStruct, val string) {
	s.val = val
}

func TestStructAsMethodArgsPointer(t *testing.T) {
	val := "struct val"
	newval := "new struct val"

	s := testStruct{val}

	noChangeTestStructVal(s, newval)
	assert.Equal(t, val, s.val)

	s = testStruct{val}

	s.noChangeVal(newval)
	assert.Equal(t, val, s.val)

	s = testStruct{val}

	changeTestStructVal(&s, newval)
	assert.Equal(t, newval, s.val)

	s = testStruct{val}

	s.changeVal(newval)
	assert.Equal(t, newval, s.val)
}
