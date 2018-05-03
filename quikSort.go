package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{9, 3, 2, 4, 8}
	kth := findKthLargest(1, nums)
	fmt.Println(nums);
	fmt.Println(kth)
}

func quickSort(nums []int) {
	if len(nums) <= 2 {
		return;
	}
	mid := nums[0]
    i , j :=0, len(nums) - 1
    for i < j {
        for i < j && mid <= nums[j] {
            j--
		}
		if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
        for i < j && nums[i] < mid {
            i++
		}
		if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
	}
	fmt.Printf("i: %d, j: %d\n", i, j)
	fmt.Println(nums)
    quickSort(nums[:i]);
    quickSort(nums[j:]);
}

func findKthLargest(k int, nums []int) int {
	mid := nums[0]
	i, j := 0, len(nums) - 1
	for {
		for i < j {
			for i < j && nums[j] <= mid {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			for i < j && mid < nums[i] {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		fmt.Println(nums)
		fmt.Println(i, j)
		if i + 1 == k {
			return nums[i]
		} else if i + 1 < k {
			i += 1
			mid = nums[i]
			j = len(nums) - 1
		} else if i + 1 > k {
			j = i - 1
			i = 0
			mid = nums[i]
		}
	}
}