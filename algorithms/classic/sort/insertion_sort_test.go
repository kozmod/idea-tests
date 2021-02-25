package sort

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	in := []int{5, 2, 4, 6, 1, 3}
	InsertionSort(in)
	assert.True(t, reflect.DeepEqual(in, []int{1, 2, 3, 4, 5, 6}))
}
