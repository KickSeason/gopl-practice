package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	wordFreq := make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordFreq[word]++ 
	}
	fmt.Println(wordFreq)
}