package main

import (
	"github.com/kozmod/idea-tests/lib/testify/main/internal"
	"testing"
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
