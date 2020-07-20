/*
插入排序
*/

package sort

// InsertionSort 插入排序，swap方式
func InsertionSort(data Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		// 寻找元素arr[i]合适的插入位置
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// TODO: 该怎么做呢？
// InsertionSort 插入排序，保存tmp变量，减少swap
// func InsertionSort(data Interface) {
// 	n := data.Len()

// 	for i := 1; i < n; i++ {
// 		// 寻找元素arr[i]合适的插入位置
// 		tmp := arr[i]
// 		var j int
// 		for j = i; j > 0 && arr[j-1] > tmp; j-- {
// 			arr[j] = arr[j-1]
// 		}
// 		arr[j] = tmp
// 	}
// }
