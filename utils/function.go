package utils

import (
	"math/rand"
	"time"
)

// GenerateRandomArray 生成有n个元素的随机数组，每个元素的随即范围为[rangeL, rangeR]
func GenerateRandomArray(n, rangeL, rangeR int) []int {
	if rangeR < rangeL {
		return []int{}
	}

	rand.Seed(time.Now().Unix())

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(rangeR-rangeL+1) + rangeL
	}
	return arr
}

// GenerateArray 生成一个有序数组, rangeL为 第一个元素的值，后续元素+1
func GenerateArray(n, rangeL int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rangeL + i
	}
	return arr
}

// GenerateNearlyOrderedArray 生成近乎有序的数组
func GenerateNearlyOrderedArray(n, swapTimes int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < swapTimes; i++ {
		posX := rand.Intn(n)
		posY := rand.Intn(n)
		arr[posX], arr[posY] = arr[posY], arr[posX]
	}
	return arr
}

// CopyArray 复制数组
func CopyArray(arr []int) []int {
	copy := make([]int, len(arr))

	for k, v := range arr {
		copy[k] = v
	}

	return copy
}

// CheckArrayIsSorted 检查数组是否已排序，从小到大
func CheckArrayIsSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
