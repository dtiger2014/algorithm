package algo_sort

import (
	"fmt"
	ds "data_struct"
)

// heapSort1, 将所有的元素依次添加到堆中, 在将所有元素从堆中依次取出来,
// 即完成了排序
// 无论是创建堆的过程, 还是从堆中依次取出元素的过程, 时间复杂度均为O(nlogn)
// 整个堆排序的整体时间复杂度为O(nlogn)
func HeapSort1(arr []int, n int) {
	mh := ds.NewMaxHeap(n)
	for i := 0; i < n; i++ {
		mh.Insert(arr[i])
	}

	for i := n - 1; i >= 0; i-- {
		e, err := mh.ExtractMax()
		if err != nil {
			fmt.Printf("HeapSort is failed, err: %v", err)
			return
		}
		arr[i] = e
	}
}

// heapSort2, 借助我们的heapify过程创建堆
// 此时, 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 
// 实践复杂度为O(nlogn)
// 堆排序的总体时间复杂度依然是O(nlogn), 但是比上述heapSort1性能更优, 
// 因为创建堆的性能更优
func HeapSort2(arr []int, n int) {
	mh := ds.MaxHeapify(arr, n)
	for i := n - 1; i >= 0; i-- {
		e, err := mh.ExtractMax()
		if err != nil {
			fmt.Printf("HeapSort is failed, err: %v", err)
			return
		}
		arr[i] = e
	}
}


// 不使用一个额外的最大堆, 直接在原数组上进行原地的堆排序
func HeapSort3(arr []int, n int) {
	// 注意，此时我们的堆是从0开始索引的
    // 从(最后一个元素的索引-1)/2开始
	// 最后一个元素的索引 = n-1
	for i := (n-1-1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}

	for i := n-1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr, i, 0)
	}
}
// 优化的shiftDown过程, 使用赋值的方式取代不断的swap,
// 该优化思想和我们之前对插入排序进行优化的思路是一致的
func shiftDown(arr []int, n, k int) {
	e := arr[k]
	for {
		if !(2*k+1 < n) {
			break
		}

		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j += 1
		}
		if e >= arr[j] {
			break
		}

		arr[k] = arr[j]
		k = j
	}

	arr[k] = e
}