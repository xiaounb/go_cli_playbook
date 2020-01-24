package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../src/command.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n') // TODO: cross-platform
		if err != nil {
			break
		}

		fmt.Println(s)
	}
}
