package main

import (
	"fmt"
)

func main() {
	nums := [...]int {1, 2, 3, 4, 5, 6}
	reverse(nums[:])
	fmt.Println(nums)
	reversep(&nums)
	fmt.Println(nums)
	res := rotate(5, nums[:])
	fmt.Println(res)
	strings := []string {"abc", "123", "123", "12345"}
	strs := removeDuplicate(strings)
	fmt.Println(strs)
}

func reverse(nums []int) {
	for i, j := 0, len(nums) -1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func reversep( p *[6]int) {
	for i, j := 0, len(p) - 1; i < j; i, j = i + 1, j - 1 {
		p[i], p[j] = p[j], p[i]
	}
}

func rotate(n int, nums []int) []int {
	if len(nums) < n {
		return nil
	}
	for i := 0; i < n; i++ {
		nums = append(nums, nums[i])
	}
	nums = nums[n:]
	return nums
}

func removeDuplicate(strings []string) []string { 
	for i := 0; i < len(strings) - 1; i++ {
		if strings[i] == strings[i + 1] {
			strings[i + 1] = strings[len(strings) - 1]
			strings = strings[:len(strings) - 1]
		}
	}
	return strings
}