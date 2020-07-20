package sort

import (
	"math/rand"
)

// QuickSort 快速排序
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	// 优化：递归到底 用 insertion
	// if l >= r {
	// 	return
	// }
	if r-l <= 15 {
		// InsertionSort(arr[l : r+1])
		return
	}

	// p := partition(arr, l, r)
	p := partition2(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// 对arr[l...r]部分进行partition操作
// 返回p，使得arr[l...p-1] < arr[p]; arr[p+1...r] > arr[p]
func partition(arr []int, l, r int) int {
	// 优化：随机拿出下标
	randIdx := rand.Intn(r-l+1) + l
	arr[l], arr[randIdx] = arr[randIdx], arr[l]

	v := arr[l]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]

	return j
}

// 双路快速排序, 为了避免数组中重复元素太多，影响效率
func partition2(arr []int, l, r int) int {
	randIdx := rand.Intn(r-l+1) + l
	arr[l], arr[randIdx] = arr[randIdx], arr[l]

	v := arr[l]

	// arr[l+1...i) <= v; arr(j...r] >= v
	i, j := l+1, r
	for {
		// 注意这里的边界，arr[i] < v, 不能是arr[i] <= v
		for {
			if !(i <= r && arr[i] < v) {
				break
			}
			i++
		}
		// 注意这里的便捷， arr[j] > v , 不能是arr[j] >= v
		for {
			if !(j >= l+1 && arr[j] > v) {
				break
			}
			j--
		}

		if i > j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// QuickSort3Ways 三路快速排序
func QuickSort3Ways(arr []int) {
	quickSort3Ways(arr, 0, len(arr)-1)
}
func quickSort3Ways(arr []int, l, r int) {
	if r-l <= 15 {
		// InsertionSort(arr[l : r+1])
		return
	}

	// partition
	randIdx := rand.Intn(r-l+1) + l
	arr[l], arr[randIdx] = arr[randIdx], arr[l]

	v := arr[l]

	lt := l     // arr[l+1...lt] < v
	gt := r + 1 // arr[gt...r] > v
	i := l + 1  // arr[lt+1...i] == v

	for {
		if !(i < gt) {
			break
		}

		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]

	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}
