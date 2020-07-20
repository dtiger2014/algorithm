package main

import (
	"fmt"
	"time"
)

/*
	leetcode 70. Climbing Stairs
	又一个楼梯，总共有n姐台阶。每一次，可以上一个台阶，也可以上两个台阶。
	问，爬上这样的一个楼梯，一共有多少不同的方法？
	- 如 n=3, 可以爬上这个楼梯的方法有：[1,1,1], [1,2], [2,1]
	- 答案：3
*/

func main() {
	n := 0
	startTime := time.Now()

	fmt.Println(climbStaris2(n))

	fmt.Println("执行时间：", time.Since(startTime))
}

func climbStaris(n int) int {
	return calcWays(n)
}

func calcWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return calcWays(n-1) + calcWays(n-2)
}

func climbStaris2(n int) int {
	if n <= 1 {
		return 1
	}
	memo := make([]int, n+1)

	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
