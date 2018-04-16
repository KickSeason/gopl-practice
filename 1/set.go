package main

import (
	"fmt"
	"os"
	"strconv"
)

func main () {
	var g int
	var fl float32 = 3.1415
	var it int = int(fl)
	fmt.Println(it)
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "please input two int\n")
		os.Exit(1)
	}
 	x, err := strconv.Atoi(os.Args[1]);
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s to int err: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[2]);
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s to int err: %v\n", os.Args[2], err)
		os.Exit(1)
	}
	g = gcd(x, y)
	fmt.Println(g)
}

func gcd (x, y int) int {
	for y != 0 {
		x, y = y, x % y
	}
	return x
}