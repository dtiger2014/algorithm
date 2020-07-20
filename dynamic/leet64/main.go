package main

import (
	"fmt"
	"math"
)

/*
64.Minimum Path Sum

给出一个m*n的矩阵，其中每一个格子包含一个非负整数。
寻找一条从左上角到右下角的路径，使得沿路的数字和最小

- 每一步只能右移或者下移
*/

func main() {
	arr := [][]int{
		{1, 1, 1},
		{2, 3, 4},
		{3, 4, 5},
	}

	fmt.Println(arr)

	res := minimumPathSum1(arr)
	fmt.Println(arr)

	fmt.Println(res)
}

// !!! 错了
func minimumPathSum(arr [][]int) []int {
	down, right := 0, 1

	memo := make([]int, 0)

	downSize, rightSize := len(arr)-1, len(arr[0])-1

	i, j := 0, 0
	for {
		if i >= downSize && j >= rightSize {
			return memo
		}

		vector := down // 向下
		if i < downSize && j < rightSize {
			// 向右
			if arr[i+1][j] > arr[i][j+1] {
				vector = right
				i++
			} else {
				j++
			}
		} else if i < downSize && j == rightSize {
			vector = right
			i++
		} else if i == downSize && j < rightSize {
			vector = down
			j++
		}
		memo = append(memo, vector)
	}
}

// 知识求和
func minimumPathSum1(arr [][]int) int {
	n, m := len(arr), len(arr[0])

	for j := 1; j < m; j++ {
		arr[0][j] += arr[0][j-1]
	}

	for i := 1; i < n; i++ {
		arr[i][0] += arr[i-1][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[i][j] += int(math.Min(float64(arr[i-1][j]), float64(arr[i][j-1])))
		}
	}

	return arr[n-1][m-1]
}
