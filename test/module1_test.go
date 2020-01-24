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

func TestJsonOutput(t *testing.T) {
	found := OpenFileAndFindString("../src/module1.txt", "go env -json")

	if !found {
		t.Errorf("'go env' should be called with the '-json' flag")
	}
}
