package main

import "testing"

func TestModule3Installation(t *testing.T) {
	found := OpenFileAndFindString("../src/module3.txt", "go get golang.org/x/tools/cmd/goimports")

	if !found {
		t.Errorf("'go get <package>' not found")
	}
}
