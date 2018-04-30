package main

import (
	"strings"
	"os"
	"fmt"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("please input filename.")
		os.Exit(1)
	}
	s := basename(os.Args[1])
	fmt.Println("basename:", s)
}

func basename(filename string) string {
	slash := strings.LastIndex(filename, "/")
	filename = filename[slash + 1:]
	if dot := strings.LastIndex(filename, ".") ; dot >= 0 {
		filename = filename[:dot]
	}
	return filename
}