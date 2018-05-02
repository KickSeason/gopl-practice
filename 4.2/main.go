package main

import (
	"os"
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		fmt.Printf("sha256: %x\n", sha256.Sum256([]byte(os.Args[1])))
		os.Exit(0)
	}
	if len(os.Args) == 3 {
		if os.Args[1] == "-a" {
			fmt.Printf("sha256: %x\n", sha256.Sum256([]byte(os.Args[2])))
			os.Exit(0)
		} else if os.Args[1] == "-b" {
			fmt.Printf("sha384: %x\n", sha512.Sum384([]byte(os.Args[2])))
			os.Exit(0)
		} else if os.Args[1] == "-c" {
			fmt.Printf("sha512: %x\n", sha512.Sum512([]byte(os.Args[2])));
			os.Exit(0)
		}
	}
}