package algo_sort

import(
	"math/rand"
	"time"
	"fmt"
)

// Generate random array, from rangL to rangeR
func GenerateRandomArray(n, rangeL, rangeR int) []int {
	arr := make([]int, 0)

	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(rangeR - rangeL) + rangeL)
	}

	return arr
}

func GenerateNearlyOrderedArray(n, swapTimes int) []int {
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	for i := 0; i < swapTimes; i++ {
		posx := rand.Intn(10000000) % n
		posy := rand.Intn(10000000) % n
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}

	return arr
}

// Copy array
func CopyArray(arr []int) []int {
	arr1 := make([]int, 0)
	arr1 = append(arr1, arr...)
	return arr1
}

// Check array is sorted
func IsSorted(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// Test sort func, and display use time.
func TestSort(sortName string, sorter func([]int, int), arr []int, n int) {
	
	startTime := time.Now()

	sorter(arr, n)

	if !IsSorted(arr, n) {
		fmt.Println("This array is not sorted!")
		return
	}

	useTime := time.Since(startTime) 

	fmt.Printf("%v : %v s", sortName, useTime)
	fmt.Println()
}