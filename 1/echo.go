package main

import (
	"fmt"
	"os"
    "strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println(os.Args[0:])
	for index, arg :=range os.Args[0:] {
		fmt.Println(index, arg)
	}
}
