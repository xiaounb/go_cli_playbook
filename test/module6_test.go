package main

import (
	"testing"
)

func TestModule6DummyTest(t *testing.T) {
	found := OpenFileAndFindString("module6.txt", "func FunctionForModule6GoDoc()")

	if !found {
		t.Errorf("goimports cannot be found")
	}

	// TODO: add more tests
}
