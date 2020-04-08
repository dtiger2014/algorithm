package main

import (
	"fmt"
	"sort"
)

type MyStr []byte

func (s MyStr) Len() int           { return len(s) }
func (s MyStr) Less(i, j int) bool { return s[i] < s[j] }
func (s MyStr) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	arr := MyStr("fasd")
	sort.Sort(arr)

	fmt.Println(string(arr))
}
