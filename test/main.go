package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := [][]int{
		// {1, 2, 3},
		// {4, 5, 6},
		// {7, 8, 9},
		// {3},
		// {2},
		// {7},
		// {9},
		// {6},
		{2, 5, 8},
		{4, 0, -1},
	}
	fmt.Println(arr)

	// fmt.Println(findDiagonalOrder(arr))

	fmt.Println("==============")

	testSort1()
}

// test sort

type sds struct {
	datas []*sd
}
type sd struct {
	id     int
	status int
}

func (arr *sds) Len() int {
	return len(arr.datas)
}
func (arr *sds) Less(i, j int) bool {
	if arr.datas[j].status < arr.datas[i].status {
		return true
	} else if arr.datas[j].status == arr.datas[i].status {
		if arr.datas[j].id > arr.datas[i].id {
			return true
		}
		return false
	} else {
		return false
	}
}
func (arr *sds) Swap(i, j int) {
	arr.datas[i], arr.datas[j] = arr.datas[j], arr.datas[i]
}

func testSort1() {
	arr := &sds{
		datas: []*sd{
			{id: 500, status: 0},
			{id: 200, status: 1},
			{id: 700, status: 2},
			{id: 300, status: 2},
			{id: 1000, status: 1},
		},
	}
	sort.Sort(arr)

	for _, v := range arr.datas {
		fmt.Printf("%+v\n", v)
	}

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

func pivotIndex(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	n := len(nums)
	suml := make([]int, n)
	suml[0] = nums[0]
	for i := 1; i < n; i++ {
		suml[i] = nums[i] + suml[i-1]
	}

	sumr := make([]int, n)
	sumr[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sumr[i] = nums[i] + sumr[i+1]
	}

	fmt.Println(suml)
	fmt.Println(sumr)

	// 判断 第一个
	if suml[0] == sumr[0] {
		return 0
	}

	for i := 1; i < n-1; i++ {
		if suml[i-1] == sumr[i+1] {
			return i
		}
	}

	// 判断最后一个
	if suml[n-1] == sumr[n-1] {
		return n - 1
	}

	return -1
}

func dominantIndex(nums []int) int {
	firstIdx, secondIdx := -1, -1
	for i, v := range nums {
		if firstIdx == -1 {
			firstIdx = i
		} else if v > nums[firstIdx] {
			secondIdx = firstIdx
			firstIdx = i
		} else {
			if secondIdx == -1 || v > nums[secondIdx] {
				secondIdx = i
			}
		}
	}

	if secondIdx == -1 || firstIdx == secondIdx || nums[firstIdx] >= nums[secondIdx]*2 {
		return firstIdx
	}
	return -1
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	digits[len(digits)-1]++
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1]++
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		ret := make([]int, 1)
		ret[0] = 1
		ret = append(ret, digits...)
		return ret
	} else {
		return digits
	}
}

func findDiagonalOrder(matrix [][]int) []int {
	ret := make([]int, 0)

	lenM := len(matrix)
	if lenM == 0 {
		return ret
	}

	lenN := len(matrix[0])
	if lenN == 0 {
		return ret
	}

	minLen := int(math.Min(float64(lenM), float64(lenN)))
	for n := 0; n < lenN+lenM-1; n++ {
		for m := 0; m < lenM; m++ {
			realm := m
			if n%2 == 0 {
				realm = lenM - 1 - m
			}

			if n < lenN && realm <= n ||
				n >= lenN && realm > n-lenN &&
					realm <= minLen+(n-lenN) {
				ret = append(ret, matrix[realm][n-realm])
			}
		}
	}
	return ret
}
