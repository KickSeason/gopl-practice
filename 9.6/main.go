package main

import (
	"fmt"
)

func main() {
	for {
		go fmt.Println(1)
		fmt.Println(0)
	}
}