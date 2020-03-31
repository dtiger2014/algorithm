package main

import (
	"fmt"
)

func main() {
	// arr := []int{3,1}
	// arr := []int{0, 1, 2, 3, 4, 5, 6}
	// arr := []int{4, 5, 6, 0, 1, 2, 3}
	// arr := []int{3, 5, 1}
	arr := []int{5, 1, 2, 3, 4}
	fmt.Println(arr)
	res := search(arr, 0)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	start, end, premid, mid := 0, n-1, 0, 0
	for end-start > 1 {
		mid = start + (end-start)>>1
		if nums[start] > nums[end] && nums[end] < nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] > nums[end] {
		start = end
	}

	end = start + n - 1
	for start <= end {
		premid = start + (end-start)>>1
		if premid < n {
			mid = premid
		} else {
			mid = premid - n
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = premid + 1
		} else {
			end = premid - 1
		}
	}
	return -1
}
