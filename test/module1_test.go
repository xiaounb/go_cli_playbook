package main

import (
	"testing"
)

func TestNakedCall(t *testing.T) {
	found := OpenFileAndFindString("../src/module1.txt", "go env")

	if !found {
		t.Errorf("'go env' not found in the command")
	}
}
