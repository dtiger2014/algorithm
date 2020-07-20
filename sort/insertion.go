/*
插入排序

- 时间复杂度：O(n^2)
- 特性：内层循环提前终止。对于有序数组，插入排序的复杂度是O(n)的。
	整体，插入排序的复杂度依然是O(n^2)的
- 插入排序：时间复杂度永远是O(n^2)。
- 使用场景：
	- 部分数据是无序的数据：日志文件，银行交易数据
*/

package sort

// InsertionSort 插入排序，swap方式
func InsertionSort(data Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		for j := i; j > 0 && data.CompareTo(j, j-1) < 0; j-- {
			data.Swap(j, j-1)
		}
	}
}

func InsertionSort2(data Interface) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		for j := i; j+1 < n && data.CompareTo(i, j+1) > 0; j++ {
			data.Swap(j, j+1)
		}
	}
}

// InsertionSort 插入排序，保存tmp变量，减少swap
func InsertionSortInts(arr []int) {
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
