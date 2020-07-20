package main

import (
	"fmt"
)

func main() {
	data := []int{24, 18, 12, 9, 16, 66, 32, 4}
	fmt.Println(Search(data, 66))
}

func Search(data []int, target int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == target {
			return i
		}
	}
	return -1
}
