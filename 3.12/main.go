package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("please input two strings.")
		os.Exit(1)
	}
	res := sameSet(os.Args[1], os.Args[2]);
	fmt.Println(res)
}

func sameSet(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, _ := range s1 {
		if !strings.Contains(s2, s1[i:i + 1]) {
			return false
		}
	}
	return true
}