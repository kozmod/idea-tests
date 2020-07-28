package big

import (
	"github.com/magiconair/properties/assert"
	"math/big"
	"testing"
)

func TestSimpleConvertBigInt(t *testing.T) {
	var smallnum, _ = new(big.Int).SetString("2188824200011112223", 10)
	num := smallnum.Uint64()
	assert.Equal(t, num, uint64(2188824200011112223))
}
