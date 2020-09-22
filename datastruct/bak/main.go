package main

import(
	sort "algo_sort"
	"fmt"
	"math/rand"
	"time"
)

func init(){
    //以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func main(){
	// algoTest1()
	// algoTest2()
	// algoTest_quick()
	algoTest_MaxHeap()
}

func algoTest_MaxHeap() {
	n := 1000000;

	// 测试1 一般性测试
	arr11 := sort.GenerateRandomArray(n, 0, n)
	arr12 := sort.CopyArray(arr11)
	arr13 := sort.CopyArray(arr11)
	arr14 := sort.CopyArray(arr11)
	arr15 := sort.CopyArray(arr11)
	arr16 := sort.CopyArray(arr11)
	arr17 := sort.CopyArray(arr11)
	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr11, n)
	sort.TestSort("Quick Sort       ", sort.QuickSort, arr12, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr13, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr14, n)
	sort.TestSort("Heap Sort 1      ", sort.HeapSort1, arr15, n)
	sort.TestSort("Heap Sort 2      ", sort.HeapSort2, arr16, n)
	sort.TestSort("Heap Sort 3      ", sort.HeapSort3, arr17, n)
	fmt.Println()

	// 测试2 测试近乎有序的数组
	arr31 := sort.GenerateNearlyOrderedArray(n, 100)
	arr32 := sort.CopyArray(arr31)
	arr33 := sort.CopyArray(arr31)
	arr34 := sort.CopyArray(arr31)
	arr35 := sort.CopyArray(arr31)
	arr36 := sort.CopyArray(arr31)
	arr37 := sort.CopyArray(arr31)
	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr31, n)
	sort.TestSort("Quick Sort       ", sort.QuickSort, arr32, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr33, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr34, n)
	sort.TestSort("Heap Sort 1      ", sort.HeapSort1, arr35, n)
	sort.TestSort("Heap Sort 2      ", sort.HeapSort2, arr36, n)
	sort.TestSort("Heap Sort 3      ", sort.HeapSort2, arr37, n)
	fmt.Println()

	// 测试3 测试存在包含大量相同元素的数组
	arr21 := sort.GenerateRandomArray(n, 0, 3)
	// arr22 := sort.CopyArray(arr21)
	arr23 := sort.CopyArray(arr21)
	arr24 := sort.CopyArray(arr21)
	arr25 := sort.CopyArray(arr21)
	arr26 := sort.CopyArray(arr21)
	arr27 := sort.CopyArray(arr21)
	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr21, n)
	// 这种情况下, 普通的QuickSort退化为O(n^2)的算法, 不做测试
	// sort.TestSort("Quick Sort", sort.QuickSort, arr22, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr23, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr24, n)
	sort.TestSort("Heap Sort 1      ", sort.HeapSort1, arr25, n)
	sort.TestSort("Heap Sort 2      ", sort.HeapSort2, arr26, n)
	sort.TestSort("Heap Sort 3      ", sort.HeapSort2, arr27, n)
}

func algoTest_quick() {

	// n = 1000000;
	// Merge Sort 2      : 323ms s
	// Quick Sort        : 102ms s
	// Quick Sort 2 Ways : 94ms s
	// Quick Sort 3 Ways : 110ms s

	// Merge Sort 2      : 183ms s
	// Quick Sort        : 58ms s
	// Quick Sort 2 Ways : 51ms s
	// Quick Sort 3 Ways : 75ms s

	// Merge Sort 2      : 250ms s
	// Quick Sort 2 Ways : 32ms s
	// Quick Sort 3 Ways : 11ms s
	
	n := 1000000;
	arr11 := sort.GenerateRandomArray(n, 0, n)
	arr12 := sort.CopyArray(arr11)
	arr13 := sort.CopyArray(arr11)
	arr14 := sort.CopyArray(arr11)

	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr11, n)
	sort.TestSort("Quick Sort       ", sort.QuickSort, arr12, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr13, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr14, n)

	fmt.Println()

	arr31 := sort.GenerateNearlyOrderedArray(n, 100)
	arr32 := sort.CopyArray(arr31)
	arr33 := sort.CopyArray(arr31)
	arr34 := sort.CopyArray(arr31)
	
	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr31, n)
	sort.TestSort("Quick Sort       ", sort.QuickSort, arr32, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr33, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr34, n)

	fmt.Println()
	
	arr21 := sort.GenerateRandomArray(n, 0, 3)
	// arr22 := sort.CopyArray(arr21)
	arr23 := sort.CopyArray(arr21)
	arr24 := sort.CopyArray(arr21)

	sort.TestSort("Merge Sort 2     ", sort.MergeSort2, arr21, n)
	// sort.TestSort("Quick Sort", sort.QuickSort, arr22, n)
	sort.TestSort("Quick Sort 2 Ways", sort.QuickSort2Ways, arr23, n)
	sort.TestSort("Quick Sort 3 Ways", sort.QuickSort3Ways, arr24, n)
}

func algoTest2() {
	n := 1000000;
	arr11 := sort.GenerateRandomArray(n, 0, n)
	arr12 := sort.CopyArray(arr11)
	arr13 := sort.CopyArray(arr11)
	arr14 := sort.CopyArray(arr11)

	sort.TestSort("Insertion Sort", sort.InsertionSort, arr11, n)
	sort.TestSort("Merge Sort", sort.MergeSort, arr12, n)
	sort.TestSort("Merge Sort 2", sort.MergeSort2, arr13, n)
	sort.TestSort("Merge Sort Bottom Up", sort.MergeSortBU, arr14, n)

	fmt.Println()

	arr31 := sort.GenerateNearlyOrderedArray(n, 100)
	arr32 := sort.CopyArray(arr31)
	arr33 := sort.CopyArray(arr31)
	arr34 := sort.CopyArray(arr31)

	sort.TestSort("Insertion Sort", sort.InsertionSort, arr31, n)
	sort.TestSort("Merge Sort", sort.MergeSort, arr32, n)
	sort.TestSort("Merge Sort 2", sort.MergeSort2, arr33, n)
	sort.TestSort("Merge Sort Bottom Up", sort.MergeSortBU, arr34, n)
}

func alogoTest1() {
	n := 20000;

	arr11 := sort.GenerateRandomArray(n, 0, n)
	arr12 := sort.CopyArray(arr11)
	arr13 := sort.CopyArray(arr11)
	arr14 := sort.CopyArray(arr11)
	arr15 := sort.CopyArray(arr11)
	arr16 := sort.CopyArray(arr11)
	arr17 := sort.CopyArray(arr11)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr11, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr12, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr13, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr14, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr15, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr16, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr17, n)
	
	fmt.Println()

	arr21 := sort.GenerateRandomArray(n, 0, 3)
	arr22 := sort.CopyArray(arr21)
	arr23 := sort.CopyArray(arr21)
	arr24 := sort.CopyArray(arr21)
	arr25 := sort.CopyArray(arr21)
	arr26 := sort.CopyArray(arr21)
	arr27 := sort.CopyArray(arr21)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr21, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr22, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr23, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr24, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr25, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr26, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr27, n)

	fmt.Println()

	arr31 := sort.GenerateNearlyOrderedArray(n, 100)
	arr32 := sort.CopyArray(arr31)
	arr33 := sort.CopyArray(arr31)
	arr34 := sort.CopyArray(arr31)
	arr35 := sort.CopyArray(arr31)
	arr36 := sort.CopyArray(arr31)
	arr37 := sort.CopyArray(arr31)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr31, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr32, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr33, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr34, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr35, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr36, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr37, n)
}