package main

import (
	"strconv"
	"fmt"
	"os"
	"./tempconv"
)

func main () {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "please input celsius.\n")
		os.Exit(1)
	}
	 c, err := strconv.Atoi(os.Args[1])
	 if err != nil {
		 fmt.Fprintf(os.Stderr, "input celsius err: %v\n", err)
		 os.Exit(1)
	 }
	 t := tempconv.CToF(tempconv.Celsius(c))
	 fmt.Println(t)
}