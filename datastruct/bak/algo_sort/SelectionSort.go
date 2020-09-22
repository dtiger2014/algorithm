package algo_sort

// Seletion sort.
func SelectionSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		// 寻找 [i, n] 区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func SelectionSort2(arr []int, n int) {
	left, right := 0, n-1

	for {
		if left >= right {
			break
		}

		minIndex, maxIndex := left, right

		if arr[minIndex] > arr[maxIndex] {
			arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
		}

		for i := left + 1; i < right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			} else if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}

		arr[left], arr[minIndex] = arr[minIndex], arr[left]
		arr[right], arr[maxIndex] = arr[maxIndex], arr[right]
		
		left++
		right--
	}
}