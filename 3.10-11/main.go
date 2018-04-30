package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("please input a number")
		os.Exit(1)
	}
	s := comma(os.Args[1])
	fmt.Println("comma", s)
}

func comma(s string) string {
	var buf bytes.Buffer
	var i int = 0
	dot := strings.LastIndex(s, ".");
	for i = 0; i < dot; i++ {
		if (dot - i) % 3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	buf.WriteString(s[dot:])
	return buf.String()
}