package main

import (
	"fmt"
	"sort"
)

/**
Students are asked to stand in non-decreasing order
of heights for an annual photo.

Return the minimum number of students not standing
in the right positions.  (This is the number of students
that must move in order for all students to be standing
in non-decreasing order of height.)



Example 1:

Input: [1,1,4,2,1,3]
Output: 3
Explanation:
Students with heights 4, 3 and the last 1 are not
standing in the right positions.


Note:

1 <= heights.length <= 100
1 <= heights[i] <= 100
*/
func main() {
	input := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(input))

	input2 := []int{1, 2, 1, 2, 1, 1, 1, 2, 1}
	fmt.Println(heightChecker(input2))
}

func heightChecker(heights []int) int {
	bak := make([]int, 0, len(heights))
	bak = append(bak, heights...)

	sort.Ints(bak)

	ret := 0
	for k, v := range heights {
		if v != bak[k] {
			ret++
		}
	}
	return ret
}
