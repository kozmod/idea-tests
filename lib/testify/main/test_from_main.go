package main

import (
	"testing"

	"github.com/kozmod/idea-tests/lib/testify/main/internal"
)

func main() {
	testSuite := []testing.InternalTest{
		{
			Name: "TestSimpleMockingFromMain",
			F:    internal.TestSimpleMockingFromMain,
		},
	}
	testing.Main(matchString, testSuite, nil, nil)
}

func matchString(a, b string) (bool, error) {
	return a == b, nil
}
