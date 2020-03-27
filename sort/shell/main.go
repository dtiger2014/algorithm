package main

import (
	"algorithm/utils"
	"fmt"
	"time"
)

func main() {
	// arr := utils.GenerateRandomArray(10000, 0, 1000)
	arr := utils.GenerateNearlyOrderedArray(100000, 10)
	// fmt.Println(arr)
	start := time.Now()
	shellSort(arr)
	fmt.Println("Shell Sort 执行时间：", time.Since(start))
	fmt.Println("Shell Sort 是否已排序: ", utils.CheckArrayIsSorted(arr))
	// fmt.Println(arr)
}

func shellSort(arr []int) {
	n := len(arr)

	// gap:分组，初始时n/2, 例如：10个元素：5,2,1
	for gap := n / 2; gap > 0; gap = gap / 2 {

		
		for i := gap; i < n; i++ {

			// 移动的方法，
			tmp := arr[i]
			var j int
			for j = i; j-gap >= 0 && arr[j-gap] > tmp; j = j - gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = tmp

			// swap方法
			// for j = i; j-gap >= 0 && arr[j-gap] > arr[j]; j = j - gap {
			// 	arr[j], arr[j-gap] = arr[j-gap], arr[j]
			// }
		}
	}
}
