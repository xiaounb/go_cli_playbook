package main

import (
	"testing"
)

func TestModule2GoFormatContent(t *testing.T) {
	numOfLines := OpenFileAndCountLines("./module2_content.go")
	if numOfLines != 8 {
		t.Errorf("it looks your 'go fmt' does not work as we expected")
	}

	// TODO: add more tests
}
