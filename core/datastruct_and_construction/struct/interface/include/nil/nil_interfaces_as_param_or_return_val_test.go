package nil

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getNilError() error {
	var e *os.PathError = nil
	return e
}

func getTrueNilError() error {
	return nil
}

func getTrueNilError2() error {
	var e error = nil
	return e
}

//An interface value is equal to nil only if both its value and dynamic type are nil.
//In the example above, getNilError() returns [nil, *os.PathError] and we compare it with [nil, nil].
//
//You can think of the interface value nil as typed, and nil without type doesn’t equal nil with type.
//If we convert nil to the correct type, the values are indeed equal.
func Test(t *testing.T) {
	shouldBeNil := getNilError()
	assert.False(t, shouldBeNil == nil)
	assert.Nil(t, shouldBeNil)

	isNil := getTrueNilError()
	assert.True(t, isNil == nil)
	assert.Nil(t, isNil)

	isNil = getTrueNilError2()
	assert.True(t, isNil == nil)
	assert.Nil(t, isNil)
}
