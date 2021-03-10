package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type eqStruct struct {
	Sval string
	Ival int
}

func TestEqualStruct_KeyByCopy(t *testing.T) {
	m := make(map[eqStruct]eqStruct)
	es := eqStruct{"a", 1}
	m[es] = es
	assert.Equal(t, m[es], es)
	assert.Equal(t, m[es], eqStruct{"a", 1})
	es.Sval = "b"
	_, ok := m[es]
	assert.False(t, ok)
}

func TestEqualStruct_KeyByPointer(t *testing.T) {
	m := make(map[*eqStruct]eqStruct)
	es := &eqStruct{"a", 1}
	m[es] = *es
	assert.Equal(t, m[es], *es)
	es.Sval = "b"
	assert.Equal(t, m[es], eqStruct{"a", 1})

	es = &eqStruct{"a", 1}
	_, ok := m[es]
	assert.False(t, ok)
}

//type notEqStruct struct {
//	Sval string
//	Slval []int
//}
//
//func TestNotEqualStructCanNotUseAsMapKey(t *testing.T) {
//	m := make(map[notEqStruct]notEqStruct)
//}
