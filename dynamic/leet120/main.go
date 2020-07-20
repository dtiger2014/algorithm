package main

import "fmt"

/*
leetcode 120. Triangle
给定一个三角形的数字阵列，选择一条自顶向下的路径，使得沿途的所有数字之和最小。
(每一步只能移动到相邻的格子中)
	[2],
   [3,4],
  [6,5,7],
 [4,1,8,3],
*/
var (
	memoIdx = []int{}
)

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	// memoIdx = make([]int, len(arr))
	// for k := range memoIdx {
	// 	memoIdx[k] = -1
	// }

	// triangle(arr, 0)
	// fmt.Println(memoIdx)

	// sum := 0
	// for k, v := range memoIdx {
	// 	sum += arr[k][v]
	// }
	// fmt.Println(sum)

	res := triangle2(arr)
	fmt.Println(res)
}

// 递归形式 + 记忆搜索
func triangle(arr [][]int, idx int) {
	if idx >= len(memoIdx) {
		return
	}

	if idx == 0 {
		memoIdx[idx] = 0
	} else {
		left, right := memoIdx[idx-1], memoIdx[idx-1]+1
		if arr[idx][left] <= arr[idx][right] {
			memoIdx[idx] = left
		} else {
			memoIdx[idx] = right
		}
	}

	triangle(arr, idx+1)

	return
}

// 动态规划 O(n)
func triangle2(arr [][]int) []int {
	memo := make([]int, len(arr))
	for k := range memo {
		memo[k] = -1
	}

	for k, v := range arr {
		if k == 0 {
			memo[k] = 0
			continue
		}

		// 先假设 左边的小？
		memo[k] = memo[k-1]	
		// 如果右边的小，则+1
		if v[memo[k-1]] > v[memo[k-1]+1] {
			memo[k]++
		}
	}

	return memo
}
