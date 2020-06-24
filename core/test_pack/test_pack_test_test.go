package test_pack_test

import (
	"testing"

	"github.com/kozmod/idea-tests/core/test_pack"
)

func TestPack(t *testing.T) {
	test_pack.BPublic()
	test_pack.BPrivate()
	Public()
	private()
}
