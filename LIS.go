package main

import (
	"fmt"
)

func main() {
	var nums = []int {2, 1, 6, 3, 5, 4, 8, 7, 9}
	cnt := lis(nums)
	fmt.Println(cnt)
}

func lis(nums []int) int {
	var max int = 0
	list := make([]int, len(nums))
	for i, _ := range list {
		list[i] = 1
	}
 	for i, v := range nums {
		for j, w := range nums[:i] {
			if w < v && list[i] < list[j] + 1 {
				list[i] = list[j] + 1
			}
		}
	}
	for _, v := range list {
		if max < v {
			max = v
		}
	}
	return max
}