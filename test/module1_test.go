package main

import (
	"testing"
)

func TestModule1NakedCall(t *testing.T) {
	found := OpenFileAndFindString("../src/module1.txt", "go env")

	if !found {
		t.Errorf("'go env' not found in the command")
	}
}

func TestModule1JsonOutput(t *testing.T) {
	found := OpenFileAndFindString("../src/module1.txt", "go env -json")

	if !found {
		t.Errorf("'go env' should be called with the '-json' flag")
	}
}
