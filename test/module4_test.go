package main

import (
	"testing"
)

func TestModule4GoFormatContent(t *testing.T) {
	expected := "	github.com/codemodus/kace"
	found := OpenFileAndFindNthString("../go.mod", 0, expected)
	if found != true {
		t.Errorf("the kace package is not found")
	}
}
