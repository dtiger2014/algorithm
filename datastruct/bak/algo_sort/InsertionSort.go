package algo_sort

// Seletion sort.
func InsertionSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		
		// 写法1
		// for j := i; j > 0; j-- {
		// 	if arr[j] < arr[j-1] {
		// 		arr[j], arr[j-1] = arr[j-1], arr[j]
		// 	} else {
		// 		break
		// 	}
		// }

		// 写法2, 写法一的简写
		// for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
		// 	arr[j], arr[j-1] = arr[j-1], arr[j]
		// }

		// 写法3 优化
		e := arr[i]
		var j int
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}

// 对arr[l...r]范围的数组进行插入排序
func InsertionSortLR(arr []int, l, r int) {
	for i := l+1; i <= r; i++ {
		e := arr[i]
		var j int
		for j = i; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}