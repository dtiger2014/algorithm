package main

import (
	"fmt"
	"math"
)

/*
279.Perfect Squares
给出一个正整数n，寻找最少的完全平方数，是他们的和为n。
- 完全平方数：1,4,9,13
- 12 = 4+4+4
- 13 = 4+9
*/

const (
	// IntMax 最大值
	IntMax = int(^uint((0)) >> 1)

	// IntMin 最小值
	IntMin = ^IntMax
)

func main() {
	n := 12

	res := perfectSquares(n)

	fmt.Println(res)
}

func perfectSquares(n int) int {
	mem := make([]int, n+1)
	for k := range mem {
		mem[k] = IntMax
	}
	mem[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; i-j*j >= 0; j++ {
			mem[i] = int(math.Min(float64(mem[i]), float64(1+mem[i-j*j])))
		}
	}
	return mem[n]
}