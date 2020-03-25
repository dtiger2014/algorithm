package main

import "fmt"

/*
!!!!!
91.Decode Ways
一个字符串，包含A-Z的字母。每一个字符可以和1-26的数字对应.
A->1, ..., Z->26 给出一个数字字符串，
问我们有多少种方法可以解析这个数字字符串？

- 如给定12，我们可以将它解析成AB（1,2）或者L（12）
- 最终返回2

*/

var (
	memo = []int{}
)

func main() {
	s := "123"

	// res := numDecodings1(s)

	res := numDecodings2(s)

	fmt.Println(res)
}

func numDecodings1(s string) int {
	memo = make([]int, len(s))
	for k := range memo {
		memo[k] = -1
	}
	return dfs(s, 0)
}

func dfs(s string, start int) int {
	if start >= len(s) {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}

	res := dfs(s, start+1)
	if start+1 < len(s) && subString(s, start, start+2) <= "26" {
		res += dfs(s, start+2)
	}

	memo[start] = res
	return res
}
func subString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

func numDecodings2(s string) int {
	n := len(s)

	if n == 1 || s[0] == '0' {

	}

	memo := make([]int, n+1)
	memo[n] = 1;
	for 
}
