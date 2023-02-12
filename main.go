package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}

	wordCount, err := wordcount(src)
	if err != nil {
		fail(err)
	}

	fmt.Println(wordCount)
}

func wordcount(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	words := strings.Fields(str)
	return len(words), nil
}

// readInput reads string from args and return it
func readInput() (string, error) {
	flag.Parse()
	src := strings.Join(flag.Args(), "")
	return strings.TrimSpace(src), nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
