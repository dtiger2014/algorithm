package main

import (
	"algorithm/heap"
	linkedlist "algorithm/linkedlist"
	"algorithm/search"
	"algorithm/sort"
	"algorithm/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// testSort()
	// testHeap()
	// testSearch()
	// testBinarySearch()

	// arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// arr := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// // OUTPUT: [1, 1, 4, 2, 1, 1, 0, 0]
	// fmt.Println(evalRPN(arr))
	// testMylinkedlist()

	fmt.Println(8^8)

	a := 0
	fmt.Println(a)
	a ^= 4
	fmt.Println(a)
	
	a ^= 8
	fmt.Println(a)
	
	a ^= 4
	fmt.Println(a)
	
	a ^= 2
	fmt.Println(a)
}

func testMylinkedlist() {
	arr := []int{1, 2, 1}

	dum := &linkedlist.ListNode{}
	head := dum
	for i := 0; i < len(arr); i++ {
		node := &linkedlist.ListNode{
			Val:  arr[i],
			Next: nil,
		}
		head.Next = node
		head = head.Next
	}
	fmt.Println(dum.Next)

	fmt.Println(linkedlist.IsPalindrome(dum.Next))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	inorderTraversalNode(root, &ret)

	return ret
}

func inorderTraversalNode(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	inorderTraversalNode(node.Left, ret)
	*ret = append(*ret, node.Val)
	inorderTraversalNode(node.Right, ret)
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for i := 0; i < len(tokens); i++ {
		char := tokens[i]

		switch char {
		case "+":
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(char)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

type MyStack struct {
	data []int
}

func NewMyStack() *MyStack {
	return &MyStack{
		data: []int{},
	}
}

func (s *MyStack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *MyStack) Pop() {
	if len(s.data) == 0 {
		return
	}

	s.data = s.data[:len(s.data)-1]
}

func (s *MyStack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func dailyTemperatures(T []int) []int {
	ret := make([]int, len(T))
	if len(T) <= 1 {
		return ret
	}

	mystack := NewMyStack()
	for i := 0; i < len(T); i++ {
		for mystack.Top() != -1 && T[i] > T[mystack.Top()] {
			ret[mystack.Top()] = i - mystack.Top()
			mystack.Pop()
		}
		mystack.Push(i)
	}
	return ret
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := []byte{s[0]}
	var char byte
	for i := 1; i < len(s); i++ {
		switch s[i] {
		case ')':
			char = '('
			break
		case ']':
			char = '['
			break
		case '}':
			char = '{'
			break
		default:
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != char {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	if len(stack) == 0 {
		return true
	}

	return false
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

func testBinarySearch() {
	n := 20000000
	arr1 := utils.GenerateArray(n, n+1)
	arr2 := utils.GenerateArray(n, 0)

	arr1 = append(arr1, arr2...)
	fmt.Println(len(arr1))

	target := 1000

	cp := time.Now()
	res := search.BinarySearchByUnSortedArray(arr1, target)
	fmt.Println("BinarySearchByUnSortedArray:", res)
	fmt.Println("BinarySearchByUnSortedArray 执行时间：", time.Since(cp))

	cp = time.Now()
	res = search.BinarySearchByUnSortedArrayNormal(arr1, target)
	fmt.Println("BinarySearchByUnSortedArrayNormal:", res)
	fmt.Println("BinarySearchByUnSortedArrayNormal 执行时间：", time.Since(cp))
}
