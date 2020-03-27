package main

import (
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	arr := utils.GenerateRandomArray(10, 0, 10)
	fmt.Println(arr)

	start := time.Now()
	mergeSort(arr)
	fmt.Println(arr)
	fmt.Println("Merge Sort 执行时间: ", time.Since(start))
	fmt.Println("Merge Sort 是否已排序：", utils.CheckArrayIsSorted(arr))
}

// mergeSort 第一种归并排序
func mergeSort(arr []int) {
	mergeSort1(arr, 0, len(arr)-1)
}
func mergeSort1(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2
	mergeSort1(arr, l, mid)
	mergeSort1(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)

	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	i, j := l, mid+1
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
