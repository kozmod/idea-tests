package _struct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type s1 struct {
	val string
}

type s2 struct {
	val int
}

func (s *s1) valS1() string {
	return s.val
}

func (s *s2) valS2() int {
	return s.val
}

type s3 struct {
	s1
	s2
}

func Test(t *testing.T) {
	s := s3{s1: s1{"a"}, s2: s2{1}}
	fmt.Println(s)
	assert.NotNil(t, s)
	assert.NotNil(t, s.s1)
	assert.NotNil(t, s.s2)
	assert.NotNil(t, s.valS1())
	assert.NotNil(t, s.valS2())
}
