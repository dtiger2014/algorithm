package hash

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	memo := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			return true
		}
		memo[v] = 1
	}
	return false
}

func singleNumber(nums []int) int {
	memo := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := memo[v]; ok {
			memo[v]++
		} else {
			memo[v] = 1
		}
	}
	for i, v := range memo {
		if v == 1 {
			return i
		}
	}
	return -1
}
