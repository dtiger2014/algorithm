package algo_sort

// 最一般的写法
func BubbleSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func BubbleSort2(arr []int, n int) {
	var swapped bool
	for {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}

		n--

		if !swapped {
			break
		}
	}
}

func BubbleSort3(arr []int, n int) {
	var newn int
	for {
		newn = 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newn = i
			}
		}

		n = newn

		if newn <= 0 {
			break
		}
	}
}