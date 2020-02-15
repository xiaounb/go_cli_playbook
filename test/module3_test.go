package main

import "testing"

func TestModule3InstallationOfgoimports(t *testing.T) {
	found := FindFileAtPath("foo", "bar")
	if !found {
		t.Errorf("goimports cannot be found")
	}
}
