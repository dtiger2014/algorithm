package sort

// MergeSort1 归并排序
func MergeSort1(arr []int) {
	mergeSortArray(arr, 0, len(arr)-1)
}
func mergeSortArray(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := l + (r-l)/2
	mergeSortArray(arr, l, mid)
	mergeSortArray(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l, mid, r int) {
	aux := make([]int, r-l+1)

	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

// MergeSort2 优化
func MergeSort2(arr []int) {
	mergeSortArray2(arr, 0, len(arr)-1)
}
func mergeSortArray2(arr []int, l, r int) {
	if r-l <= 15 {
		InsertionSort(arr[l : r+1])
		return
	}

	mid := l + (r-l)/2
	mergeSortArray2(arr, l, mid)
	mergeSortArray2(arr, mid+1, r)

	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

// MergeSortBU 归并排序自底向上进行归并
func MergeSortBU(arr []int) {
	n := len(arr)

	for sz := 1; sz < n; sz += sz{
		for i := 0; i < n-sz; i += sz+sz {
			
		}
	}
}
