package main

import "fmt"

/*
62.Unique Paths
Bloomberg
又一个机器人，从一个m*n的矩阵的左上角出发，要达到这个矩阵的右下角。
机器人每次只能向右或者向下进行。
文艺工有多少中不同的路径？
*/

func main() {
	m, n := 3, 3

	res := uniquePaths(m, n)

	fmt.Println(res)
}

func uniquePaths(m, n int) int {

	return 0
}

func findWays(m, n int, vector) int {

}