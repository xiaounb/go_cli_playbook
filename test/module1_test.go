package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestNakedCall(t *testing.T) {
	found := openFileAndFindString("go env")

	if found == false {
		t.Errorf("Not found 'go env' in the command file")
	}
}

func findNthCommand() string {
	return "Got"
}

func openFileAndFindString(expected string) bool {
	f, err := os.Open("../src/module1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		trimmed := strings.Trim(t, " ")
		if trimmed == "" {
			continue
		}
		if trimmed == expected {
			return true
		}
	}

	return false
}
