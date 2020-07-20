/*
选择排序
*/

package sort

// SelectionSort 选择排序
func SelectionSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		// 寻找[i, n)区间里的最小值
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data.CompareTo(j, minIdx) < 0 {
				minIdx = j
			}
		}
		data.Swap(i, minIdx)
	}
}

// SelectionSort2 选择排序 实现2
func SelectionSort2(data Interface) {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		// 确保 [0, i] 区间最最大值值
		maxIdx := i
		for j := i - 1; j >= 0; j-- {
			if data.CompareTo(j, maxIdx) > 0 {
				maxIdx = j
			}
		}
		data.Swap(i, maxIdx)
	}
}
