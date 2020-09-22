package util


func Sum(arr []int) int {
	return sum(arr, 0)
}

func sum(arr []int, i int) int {
	if i == len(arr) {
		return 0
	}
	return arr[i] + sum(arr, i + 1)
}