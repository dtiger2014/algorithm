package main

import (
	"fmt"
)

/**
Given an array of integers A sorted in non-decreasing order,
return an array of the squares of each number,
also in sorted non-decreasing order.

Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/

func main() {
	input1 := []int{-4, -1, 0, 3, 10}
	output1 := sortedSquares(input1)
	fmt.Println(output1)

	input2 := []int{-7, -3, 2, 3, 11}
	output2 := sortedSquares(input2)
	fmt.Println(output2)
}

func sortedSquares(A []int) []int {
	ret := make([]int, len(A))
	l, r, p := 0, len(A), len(A)-1

	for {
		if l >= r {
			return ret
		}

		x, y := A[l]*A[l], A[r-1]*A[r-1]
		if x >= y {
			ret[p] = x
			p--
			l++
		} else {
			ret[p] = y
			p--
			r--
		}
	}
}
