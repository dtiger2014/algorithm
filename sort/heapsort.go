package sort

import (
	"algorithm/heap"
)

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 直接使用Hepify方法，创建堆
	heap := heap.NewMaxHeapByArray(arr)

	for i := heap.Size() - 1; i >= 0; i-- {
		arr[i], _ = heap.ExtractMax()
	}
}

// HeapSort2 堆排序
func HeapSort2(arr []int) {
	heap := heap.NewMaxHeap()

	for _, v := range arr {
		heap.Insert(v)
	}

	for i := heap.Size() - 1; i >= 0; i-- {
		arr[i], _ = heap.ExtractMax()
	}
}

// HeapSort3 堆排序，原地排序。不适用一个额外的最大堆，直接在元数组上进行原地的堆排序
func HeapSort3(arr []int) {
	// 此时堆是从0开始索引的
	// 从（最后一个元素-1)/2开始
	// 最后一个元素的索引 = n - 1
	for i := (len(arr) - 1 - 1) / 2; i >= 0; i-- {
		shiftdown2(arr, len(arr), i)
	}

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftdown2(arr, i, 0)
	}
}

// 原地shiftdown 过程
func shiftdown(arr []int, n, k int) {
	for {
		if !(2*k+1 < n) {
			break
		}
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j++
		}
		if arr[k] >= arr[j] {
			break
		}

		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}

// 优化的shiftdown过程，使用赋值的方式取代不断的swap
// 该优化思想和插入排序优化的思路是一致的
func shiftdown2(arr []int, n, k int) {
	e := arr[k]
	for {
		if !(2*k+1 < n) {
			break
		}
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j++
		}
		if e >= arr[j] {
			break
		}

		arr[k] = arr[j]
		k = j
	}
	arr[k] = e
}
