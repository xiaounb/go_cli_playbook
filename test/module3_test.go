package main

import (
	"testing"
)

func TestModule3GoFormatContent(t *testing.T) {
	expected := "	github.com/nathan-osman/go-sunrise"
	found := OpenFileAndFindNthString("../go.mod", 0, expected)
	if found != true {
		t.Errorf("the go-sunrise package is not found")
	}
}
