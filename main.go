package main

import (
	"algorithm/heap"
	"algorithm/search"
	"algorithm/sort"
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	// testSort()
	// testHeap()
	// testSearch()
}

func testHeap() {
	n := 10

	arr := utils.GenerateRandomArray(n, 0, n*n)

	fmt.Println(arr)
	heap := heap.NewMaxHeap()
	for _, v := range arr {
		heap.Insert(v)
	}

	for range arr {
		v, err := heap.ExtractMax()
		fmt.Println("Heap extractMax:", v, err)
	}
}

func testSort() {
	n := 1000000 // 数组元素数量
	// swaptime := 10
	arr1 := utils.GenerateRandomArray(n, 0, n)
	// arr1 = utils.GenerateNearlyOrderedArray(n, swaptime)

	arr2 := utils.CopyArray(arr1)
	arr3 := utils.CopyArray(arr1)
	arr4 := utils.CopyArray(arr1)
	arr5 := utils.CopyArray(arr1)
	arr6 := utils.CopyArray(arr1)
	arr7 := utils.CopyArray(arr1)
	arr8 := utils.CopyArray(arr1)
	arr9 := utils.CopyArray(arr1)
	arr10 := utils.CopyArray(arr1)
	arr11 := utils.CopyArray(arr1)
	arr12 := utils.CopyArray(arr1)

	var title string
	var cpTime time.Time

	if n <= 10000 {
		title = "Bubble"
		cpTime = time.Now()
		sort.BubbleSort(arr1)
		fmt.Println(title, "执行时间：", time.Since(cpTime))
		fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr1))

		title = "Selection"
		cpTime = time.Now()
		sort.SelectionSort(arr2)
		fmt.Println(title, "执行时间：", time.Since(cpTime))
		fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr2))

		title = "Insertion"
		cpTime = time.Now()
		sort.InsertionSort(arr3)
		fmt.Println(title, "执行时间：", time.Since(cpTime))
		fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr3))
	}

	title = "Shell"
	cpTime = time.Now()
	sort.ShellSort(arr4)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr4))

	title = "MergeSort1"
	cpTime = time.Now()
	sort.MergeSort1(arr5)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr5))

	title = "MergeSort2"
	cpTime = time.Now()
	sort.MergeSort2(arr6)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr6))

	title = "MergeSortBU"
	cpTime = time.Now()
	sort.MergeSortBU(arr7)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr7))

	title = "QuickSort"
	cpTime = time.Now()
	sort.QuickSort(arr8)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr8))

	title = "QuickSort3Ways"
	cpTime = time.Now()
	sort.QuickSort3Ways(arr9)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr9))

	title = "HeapSort2"
	cpTime = time.Now()
	sort.HeapSort2(arr10)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr10))

	title = "HeapSort"
	cpTime = time.Now()
	sort.HeapSort(arr11)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr11))

	title = "HeapSort3"
	cpTime = time.Now()
	sort.HeapSort3(arr12)
	fmt.Println(title, "执行时间：", time.Since(cpTime))
	fmt.Println(title, "是否排序排序：", utils.CheckArrayIsSorted(arr12))

}

func testSearch() {
	arr := utils.GenerateRandomArray(5, 0, 100)
	sort.QuickSort3Ways(arr)
	fmt.Println(arr)

	floor := search.BinarySearchFloor(arr, arr[2])
	fmt.Println(floor)
	ceil := search.BinarySearchCeil(arr, arr[2])
	fmt.Println(ceil)
	// index := search.BinarySearch(arr, 200)
	// fmt.Println(index, arr[index])
}