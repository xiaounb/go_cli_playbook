package main

import (
	"os"
	"path"
	"testing"
)

func TestModule5InstallationOfgoimports(t *testing.T) {
	binPath := path.Join(os.Getenv("GOPATH"), "bin")
	found := FindFileAtPath(binPath, "goimports")

	if !found {
		t.Errorf("goimports cannot be found")
	}
}
