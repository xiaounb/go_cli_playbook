package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		fmt.Println(variable[0], "=>", variable[1])
	}

	fmt.Println(os.Getenv("GOPATH"))
}
