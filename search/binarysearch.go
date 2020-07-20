package search

// 二分查找法，在有序属组arr中，查找target
// 如果找到target，返回相应的索引index
// 如果没有找到target，返回-1

// BinarySearch 二分查找法
func BinarySearch(arr []int, target int) int {

	l, r := 0, len(arr)-1
	for {
		if !(l <= r) {
			return -1
		}

		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}

// 二分查找法, 在有序数组arr中, 查找target
// 如果找到target, 返回第一个target相应的索引index
// 如果没有找到target, 返回比target小的最大值相应的索引, 如果这个最大值有多个, 返回最大索引
// 如果这个target比整个数组的最小元素值还要小, 则不存在这个target的floor值, 返回-1

// BinarySearchFloor 二分查找法变种：floor
func BinarySearchFloor(arr []int, target int) int {
	l, r := -1, len(arr)-1
	for {
		if !(l < r) {
			goto JUMP
		}

		mid := l + (r-l+1)/2
		if arr[mid] >= target {
			r = mid - 1
		} else {
			l = mid
		}
	}
JUMP:
	if l+1 >= 0 && arr[l+1] == target {
		return l + 1
	}
	return l
}

// 二分查找法, 在有序数组arr中, 查找target
// 如果找到target, 返回最后一个target相应的索引index
// 如果没有找到target, 返回比target大的最小值相应的索引, 如果这个最小值有多个, 返回最小的索引
// 如果这个target比整个数组的最大元素值还要大, 则不存在这个target的ceil值, 返回整个数组元素个数n

// BinarySearchCeil 二分查找法变种：ceil
func BinarySearchCeil(arr []int, target int) int {
	l, r := 0, len(arr)
	for {
		if !(l < r) {
			goto JUMP
		}

		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
JUMP:
	if r-1 >= 0 && arr[r-1] == target {
		return r - 1
	}
	return r
}

// BinarySearchByUnSortedArrayNormal 循环查找法，
func BinarySearchByUnSortedArrayNormal(nums []int, target int)int {
	if len(nums) == 0 {
		return -1
	}
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}
// BinarySearchByUnSortedArray 二分查找法，在数组中间截断不满足顺序。例如:[3,4,5,0,1,2]
func BinarySearchByUnSortedArray(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	start, end, premid, mid := 0, n-1, 0, 0
	for end-start > 1 {
		mid = start + (end-start)>>1
		if nums[start] > nums[end] && nums[end] < nums[mid] {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] > nums[end] {
		start = end
	}

	end = start + n - 1
	for start <= end {
		premid = start + (end-start)>>1
		if premid < n {
			mid = premid
		} else {
			mid = premid - n
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = premid + 1
		} else {
			end = premid - 1
		}
	}
	return -1
}
