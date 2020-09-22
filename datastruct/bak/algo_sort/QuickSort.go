// 比较Merge Sort和Quick Sort两种排序算法的性能效率
// 两种排序算法虽然都是O(nlogn)级别的, 但是Quick Sort算法有常数级的优势
// Quick Sort要比Merge Sort快, 即使我们对Merge Sort进行了优化

// 测试存在包含大量相同元素的数组
// 使用双快速排序后, 我们的快速排序算法可以轻松的处理包含大量元素的数组

package algo_sort

import(
	// "fmt"
	"math/rand"
)


func QuickSort(arr []int, n int) {
	quickSort(arr, 0, n-1)
}

func quickSort(arr []int, l, r int) {
	// 优化：对于小规模数组，使用插入排序进行优化
	if r - l <= 15 {
		InsertionSortLR(arr, l, r)
		return
	}

	// p := partition(arr, l, r)
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// 对arr[l...r]部分进行partition操作
func partition(arr []int, l, r int) int {

	// 随即在arr[l...r]的范围中，选择一个数值作为标定点pivot
	randIndex := rand.Intn(len(arr)) % (r - l + 1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]

	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[l], arr[j] = arr[j], arr[l]

	return j
}

func QuickSort2Ways(arr []int, n int) {
	quickSort2(arr, 0, n-1)
}

func quickSort2(arr []int, l, r int) {
	// 优化：对于小规模数组，使用插入排序进行优化
	if r - l <= 15 {
		InsertionSortLR(arr, l, r)
		return
	}

	p := partition2(arr, l, r)
	quickSort2(arr, l, p-1)
	quickSort2(arr, p+1, r)
}

// 双路快速排序的partition
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func partition2(arr []int, l, r int) int {
	// 随即在arr[l...r]的范围中，选择一个数值作为标定点pivot
	randIndex := rand.Intn(len(arr)) % (r - l + 1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]

	// arr[l+1...i] <= v; arr(j....r] >= v
	i, j := l+1, r
	for {
		for {
			if !(i <= r && arr[i] < v) {
				goto CHECKJ
			}
			i++
		}
	CHECKJ:
		for {
			if !(j >= l+1 && arr[j] > v) {
				goto SWAP
			}
			j--
		}

	SWAP:
		if i > j {
			break;
		}

		arr[i], arr[j] = arr[j], arr[i]

		i++
		j--
	}

	arr[l], arr[j] = arr[j], arr[l]
	
	return j
}


// 递归的三鹿快速排序算法
func QuickSort3Ways(arr []int, n int) {
	quickSort3Ways(arr, 0, n-1)
}

func quickSort3Ways(arr []int, l, r int) {
	// 对于小规模数组，使用插入排序进行优化
	if r - l <= 15 {
		InsertionSortLR(arr, l, r)
		return
	}

	// 随即在arr[l...r]的范围中，选择一个数值作为标定点pivot
	randIndex := rand.Intn(len(arr)) % (r - l + 1) + l
	arr[l], arr[randIndex] = arr[randIndex], arr[l]

	v := arr[l]

	// arr[l+1...lt] < v, arr[gt....r] > v, arr[lt+1....i] == v
	lt, gt, i := l, r+1, l+1
	for {
		if !(i < gt) {
			break
		}

		if arr[i] < v {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > v {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else { // arr[i] == v
			i++
		}
	} 

	arr[l], arr[lt] = arr[lt], arr[l]

	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}