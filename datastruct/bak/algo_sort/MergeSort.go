// Merge Sort BU 也是一个O(nlogn)复杂度的算法，虽然只使用两重for循环
// 所以，Merge Sort BU也可以在1秒之内轻松处理100万数量级的数据
// 注意：不要轻易根据循环层数来判断算法的复杂度，Merge Sort BU就是一个反例

package algo_sort

import(
	// "fmt"
)

// Merge Sort 1
func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return;
	}
	
	mid := (l + r) / 2
	
	mergeSort(arr, l, mid)
	mergeSort(arr, mid + 1, r)
	merge(arr, l, mid, r)
}

// Merge Sort 2
func MergeSort2(arr []int, n int) {
	mergeSort2(arr, 0, n-1)
}

func mergeSort2(arr []int, l, r int) {
	// 优化：对于小规模数组，使用插入排序
	if r - l <= 15 {
		InsertionSortLR(arr, l, r)
		return
	}

	mid := (l + r) / 2
	mergeSort2(arr, l, mid)
	mergeSort2(arr, mid+1, r)

	// 优化：对于arr[mid] <= arr[mid+1]的情况，不进行merge
	// 对于近乎有序的数组非常有效，但是对于一般情况，有一定的性能损失
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

// 使用自底向上的归并排序算法
func MergeSortBU(arr []int, n int) {
	// Merge Sort Bottom Up 无优化版
	// for sz := 1; sz < n; sz += sz {
	// 	for i := 0; i < n - sz; i += sz+sz {
	// 		// 对 arr[i.....i+sz-1] 和 arr[i+sz....i+2*sz-1] 进行归并
	// 		merge(arr, i, i+sz-1, min(i+sz+sz-1, n-1));
	// 	}
	// }

	// ======

	// Merge Sort Bottom Up 优化
	// 对于小数组，使用插入排序优化
	for i := 0; i < n; i += 16 {
		InsertionSortLR(arr, i, min(i+15, n-1))
	}

	for sz := 16; sz < n; sz += sz {
		for i := 0; i < n - sz; i += sz+sz {
			// 对于arr[mid] <= arr[mid+1]的情况，不进行merge
			if arr[i+sz-1] > arr[i+sz] {
				merge(arr, i, i+sz-1, min(i+sz+sz-1, n-1))
			}
		}
	}
}

// merge
func merge(arr []int, l, mid, r int) {
	aux := make([]int, 0)

	for i := l; i <= r; i++ {
		aux = append(aux, arr[i])
	}

	i, j := l, mid + 1
	for k := l; k <= r; k++ {
		if i > mid { 
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
} 

// common: min
func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}