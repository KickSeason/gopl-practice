package main

import (
	"os"
	"fmt"
	"crypto/sha256"
)

func main() {
	if (len(os.Args) < 3) {
		fmt.Println("please input two string.")
		os.Exit(1)
	}
	s1 := sha256.Sum256([]byte(os.Args[1]))
	fmt.Printf("s1: %x\n", s1)
	s2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("s2: %x\n", s2)
	cnt := DiffBitCount(s1, s2)
	fmt.Println("diff cunt:", cnt)
}

func PopCountAnd(x uint64) int {
	var popcount int
	for  0 < x {
		x = x & (x - 1)
		popcount++
	}
	return popcount
}

func DiffBitCount(s1, s2 [32]byte) int {
	var cnt int
	for i, v := range s1 {
		cnt += PopCountAnd(uint64(v ^ s2[i]))
	}
	return cnt
}