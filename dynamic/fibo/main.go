package main

import (
	"fmt"
	"time"
)

var (
	num  = 0
	memo = []int{}
)

func main() {

	n := 100
	startTime := time.Now()

	memo := make([]int, n+1)
	for k := range memo {
		memo[k] = -1
	}

	fmt.Println(fib2(n))

	fmt.Println("执行时间：", time.Since(startTime))
	fmt.Println("计算次数: ", num)
}

/*
菲波那切数列 Fibonacci Sequence
F(0) = 1, F(1)=1, F(n)=F(n-1)+F(n-2)
*/
func fib(n int) int {
	num++
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// 数组记录 计算的结果，记忆话搜索
// 减少计算次数
// 自上向下的解决问题
func fib1(n int) int {
	num++
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if memo[n] == -1 {
		memo[n] = fib1(n-1) + fib1(n-2)
	}

	return memo[n]
}


// 自下向上的解决问题
func fib2(n int) int {
	num++
	memo := make([]int, n+1)

	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
