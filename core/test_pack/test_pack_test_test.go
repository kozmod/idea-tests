package test_pack_test

import (
	"github.com/kozmod/idea-tests/core/test_pack"
	"testing"
)

func TestPack(t *testing.T) {
	test_pack.BPublic()
	test_pack.BPrivate()
	Public()
	private()
}
