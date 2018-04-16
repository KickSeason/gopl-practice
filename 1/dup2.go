package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "open %s failed, err: %v.\n", filename, err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}
	for line, names := range counts {
		count := len(strings.Split(names, " "))
		if 2 < count {
			fmt.Printf("%s\t%s\n", names, line)
		}
	}
}

func countLines(file *os.File, counts map[string]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		text := input.Text()
		counts[text] += " " + file.Name();
	}
}
