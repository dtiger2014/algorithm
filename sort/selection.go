package sort

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// 寻找[i, n)区间里的最小值
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}