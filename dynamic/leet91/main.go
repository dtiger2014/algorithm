package main

import (
	"fmt"
	"strconv"
)

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
	s := "12"

	// res := numDecodings1(s)

	res, op := numDecodings3(s)
	fmt.Println(res)
	fmt.Println(op)
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

// 数字字符串，最小的问题是 可以从最后一个字符开始
func numDecodings2(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}

	if n == 1 {
		return 1
	}

	memo := make([]int, n+1)
	memo[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			memo[i] = memo[i+1]
			if i+1 < n && subString(s, i, 2) <= "26" {
				memo[i] += memo[i+2]
			}
		}
	}

	return memo[0]
}

// 输出所有结果
func numDecodings3(s string) (int, []string) {
	op := make([]string, 0)

	n := len(s)
	if s[0] == '0' {
		return 0, op
	}

	if n == 1 {
		op = append(op, getChar(s))
		return 1, op
	}

	memo := make([]int, n+1)
	memo[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			
			memo[i] = memo[i+1]
			
			if i+1 < n && subString(s, i, 2) <= "26" {
				memo[i] += memo[i+2]
			}
		}
	}

	return memo[0], op
}

func getChar(s string) string {
	num, _ := strconv.Atoi(s)
	res := string(num - 1 + 'A')
	return res
}

func saveOP(op []string, str string, isNew bool) {
	n := len(op)
	if isNew {
		for i := 0; i < n; i++ {
			op = append(op, str+op[i])
		}
	} else {
		for k, v := range op {
			op[k] = str + v
		}
	}
	
	fmt.Println(op)
}
