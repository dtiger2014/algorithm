package main

import (
	"algorithm/sort"
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	n := 1000000 // 数组元素数量
	swaptime := 10
	arr1 := utils.GenerateRandomArray(n, 0, n)
	arr1 = utils.GenerateNearlyOrderedArray(n, swaptime)
	
	arr2 := utils.CopyArray(arr1)
	arr3 := utils.CopyArray(arr1)
	arr4 := utils.CopyArray(arr1)
	arr5 := utils.CopyArray(arr1)
	arr6 := utils.CopyArray(arr1)

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
}
