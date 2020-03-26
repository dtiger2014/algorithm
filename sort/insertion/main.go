package main

import (
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	// fmt.Println(utils.GenerateNearlyOrderedArray(10,2))

	// arr := utils.GenerateRandomArray(50000, 0, 3)
	arr := utils.GenerateNearlyOrderedArray(50000, 10)
	arr2 := utils.CopyArray(arr)
	start := time.Now()
	insertionSort(arr)
	fmt.Println("Insertion Sort1 执行时间：", time.Since(start))

	start = time.Now()
	insertionSort2(arr2)
	fmt.Println("Insertion Sort2 执行时间：", time.Since(start))

	// fmt.Println("Insertion Sort 是否已排序：", utils.CheckArrayIsSorted(arr))
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// 优化：减少 swap
func insertionSort2(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		tmp := arr[i]
		var j int
		for j = i; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}
