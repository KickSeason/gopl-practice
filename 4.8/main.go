package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	dict := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		fmt.Println(word)
		dict[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	for word, value := range dict {
		fmt.Println(word, ":", value)
	}
}