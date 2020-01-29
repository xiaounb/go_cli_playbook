package main

import (
	"testing"
)

func TestModule2NakedCall(t *testing.T) {
	found := OpenFileAndFindString("../src/module2.txt", "gofmt common.go")

	if !found {
		t.Errorf("'gofmt <filename>' not found")
	}
}

func TestModule2DefaultParams(t *testing.T) {
	found := OpenFileAndFindString("../src/module2.txt", "gofmt -l -w common.go")

	if !found {
		t.Errorf("See 'go help fmt' for more details")
	}
}
