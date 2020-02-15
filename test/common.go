package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

// FindFileAtPath returns if theFilename is found at thePath
func FindFileAtPath(thePath string, theFilename string) bool {
	if _, err := os.Stat(path.Join(thePath, theFilename)); os.IsNotExist(err) {
		return false
	}
	return true
}

// OpenFileAndCountLines opens a file and returns number of lines in the file
func OpenFileAndCountLines(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	number := 0
	for scanner.Scan() {
		number++
	}

	return number
}

// OpenFileAndFindString opens a file and return if the given string is found or not
func OpenFileAndFindString(filename string, expected string) bool {
	f, err := os.Open(filename)
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
