package  main

import (
	"./popcount"
	"os"
	"fmt"
	"strconv"
)

func main() {
	var num uint64
	if (len(os.Args) < 2) {
		fmt.Println("please input uint64!")
		os.Exit(1)
	}
	ii, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid number.");
		os.Exit(1);
	}
	num = uint64(ii)
	pc := popcount.PopCount(num)
	fmt.Println("positive count: ", pc)
	pc = popcount.PopCountBit(num)
	fmt.Println("positive count: ", pc)
	pc = popcount.PopCountAnd(num)
	fmt.Println("positive count: ", pc)
}