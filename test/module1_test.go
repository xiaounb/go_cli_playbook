package main

import (
	"os"
	"path"
	"testing"
)

func TestModule1CheckEnv(t *testing.T) {
	actual := os.Getenv("GOPATH")
	expected := path.Join(os.Getenv("HOME"), "go")

	if actual != expected {
		t.Errorf("environment variable GOPATH not set properly")
	}
}
