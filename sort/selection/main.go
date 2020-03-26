package main

import (
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	arr := utils.GenerateRandomArray(10000, 0, 100000)
	start := time.Now()
	selectionSort(arr)
	fmt.Println("Selection Sort 执行时间：", time.Since(start))
	fmt.Println("Selection Sort 是否以排序：", utils.CheckArrayIsSorted(arr))
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// 寻找[i, n)区间里的最小值
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}