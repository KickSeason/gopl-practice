package main

import (
	"fmt"
)
type tree struct {
	value		int
	left, right	*tree
}

func main() {
	nums := []int {2, 5, 8, 3, 0, 9, 1}
	result := treeSort(nums)
	fmt.Println(result)
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value =  value;
		return root
	}
	if root.value < value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func appendValues(nums []int, root *tree) []int {
	if root != nil {
		nums = appendValues(nums, root.left)
		nums = append(nums, root.value)
		nums = appendValues(nums, root.right)
	}
	return nums
}

func treeSort(nums []int) []int {
	var root *tree
	for _, val := range nums {
		root = add(root, val)
	}
	nums = appendValues(nums[:0], root)
	return nums
}