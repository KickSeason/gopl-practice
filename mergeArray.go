package main

import (
	"fmt"
)

func main() {
	var A []int = []int {1, 5}
	var B = []int {2, 3}
	C := mergeSortedArray(A, B)
	fmt.Println(C)
}

func mergeSortedArray (A []int, B []int) []int {
    // write your code here
    var merg []int
    ia, ib := 0, 0
    for ia < len(A) || ib < len(B) {
        if len(B) <= ib || (ia < len(A) && A[ia] <= B[ib])  {
            merg = append(merg, A[ia])
            ia++
        } else {
            merg = append(merg, B[ib])
            ib++
		}
    }
    return merg;
}