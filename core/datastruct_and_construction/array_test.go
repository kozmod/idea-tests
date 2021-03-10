package datastruct_and_construction_test

import (
	"fmt"
	"testing"
)

func TestArrayVsSlice(t *testing.T) {
	t.Skip()
	//a := [3]string{"a", "b", "c","d"}
	s := []string{"a", "b", "c", "d"}
	//a = append(a, "a")
	s = append(s, "a")
	//fmt.Println(len(a), len(a))
	fmt.Println(cap(s), cap(s))
}
