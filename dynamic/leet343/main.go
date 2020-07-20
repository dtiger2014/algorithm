package main

import (
	"fmt"
	"math"
)

/*
343.Integer Break
给一个正数n，可以将其分割成多个数字的和，若要让这些数字的乘机最大，
求分隔的方法（至少要分成两个数）。
算法返回这个最大的乘积

- 如 n=2, 则返回1（2=1+1）
- 如 n=10，则返回36 （10=3+3+4）

暴力解法：回溯遍历讲一个数做分隔的所有可能性. O（2^n）

最优子结构：通过求子问题的最优解，可以获得原问题的最优解
*/

var (
	memo = []int{}
)

func main() {
	n := 11

	res := integerBreak2(n)

	fmt.Println(res)
}

// 最优子结构
func integerBreak1(n int) int {
	if n <= 2 {
		panic("err")
	}
	memo = make([]int, n+1)
	for v := range memo {
		memo[v] = -1
	}
	return breakInteger(n)
}

// 将n进行分割（至少分割两部分），可以获得的最大乘积
func breakInteger(n int) int {
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	res := -1
	for i := 1; i <= n-1; i++ {
		// i + (n-i) = n
		// i*(n-i) : 查看 如果只分2个是否为最大？
		// i*breakInteger(n-i): 查看 分成2个以上，是否为最大？
		res = max3(res, i*(n-i), i*breakInteger(n-i))
	}
	memo[n] = res
	return res
}

func max3(a, b, c int) int {
	return int(math.Max(float64(c), math.Max(float64(a), float64(b))))
}

// 动态规划方法

func integerBreak2(n int) int {
	if n < 2 {
		panic("err")
	}

	memo := make([]int, n+1)
	for k := range memo {
		memo[k] = -1
	}

	memo[1] = 1
	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			// j + (i - j)
			memo[i] = max3(memo[i], j*(i-j), j*memo[i-j])
		}
	}

	return memo[n]
}
