package array

import (
	"errors"
	"fmt"
	"golang-datastruct/util/common"
)

// ArrayData : 动态数组 struct
type ArrayData struct {
	data []interface{}
	size int
}

// New : 创建初始化，并返回Array
func New() *ArrayData {
	return &ArrayData{
		data: make([]interface{}, common.DefaultCap, common.DefaultCap),
		size: 0,
	}
}

// GetCapacity : 获取Capacity
func (arr *ArrayData) GetCapacity() int {
	return cap(arr.data)
}

// GetSize : 获取Size
func (arr *ArrayData) GetSize() int {
	return arr.size
}

// IsEmpty : 判断是否为空
func (arr *ArrayData) IsEmpty() bool {
	return arr.size == 0
}

// Add : 添加元素
func (arr *ArrayData) Add(index int, e interface{}) error {
	if index < 0 || index > arr.size {
		return errors.New("Add failed. Require index >= 0 and index <= size")
	}

	if arr.size == cap(arr.data) {
		arr.resize(2 * cap(arr.data))
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}

	arr.data[index] = e
	arr.size++

	return nil
}

// AddLast : 添加元素 Last
func (arr *ArrayData) AddLast(e interface{}) {
	arr.Add(arr.size, e)
}

// AddFirst : 添加元素 First
func (arr *ArrayData) AddFirst(e interface{}) {
	arr.Add(0, e)
}

// Swap：
func (arr *ArrayData) Swap(i, j int) error {
	if i < 0 ||
		i >= arr.size ||
		j < 0 ||
		j >= arr.size {
		return errors.New("Index is illegal")
	}

	arr.data[i], arr.data[j] = arr.data[j], arr.data[i]
	return nil
}

// Get : 获取元素
func (arr *ArrayData) Get(index int) (interface{}, error) {
	if index < 0 || index > arr.size {
		return nil, errors.New("Get failed. Index is illegal")
	}

	return arr.data[index], nil
}

// GetFirst : 获取元素 First
func (arr *ArrayData) GetFirst() (interface{}, error) {
	return arr.Get(0)
}

// GetLast : 获取元素 Last
func (arr *ArrayData) GetLast() (interface{}, error) {
	return arr.Get(arr.size - 1)
}

// Set : 设置元素
func (arr *ArrayData) Set(index int, e interface{}) error {
	if index < 0 || index > arr.size {
		return errors.New("Set failed. Index is illegal")
	}

	arr.data[index] = e

	return nil
}

// Contains : 判断数组是否包含目标元素
func (arr *ArrayData) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// Find : 返回目标元素下标，无 返回 -1
func (arr *ArrayData) Find(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

// Remove : 删除元素
func (arr *ArrayData) Remove(index int) (interface{}, error) {
	if index < 0 || index >= arr.size {
		return nil, errors.New("Remove failed. Index is illegal")
	}

	ret := arr.data[index]

	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	arr.data[arr.size] = nil

	if arr.size == cap(arr.data)/4 && cap(arr.data)/2 != 0 {
		arr.resize(cap(arr.data) / 2)
	}

	return ret, nil
}

// RemoveFirst : 删除元素 First
func (arr *ArrayData) RemoveFirst() (interface{}, error) {
	return arr.Remove(0)
}

// RemoveLast : 删除元素 Last
func (arr *ArrayData) RemoveLast() (interface{}, error) {
	return arr.Remove(arr.size - 1)
}

// RemoveElement : 删除目标元素
func (arr *ArrayData) RemoveElement(e interface{}) {
	index := arr.Find(e)
	if index != 1 {
		_, _ = arr.Remove(index)
	}
}

// String : 格式化输出
func (arr *ArrayData) String() string {
	str := fmt.Sprintf("Array: size = %d , capacity = %d\n", arr.size, cap(arr.data))
	str += "["
	for i := 0; i < arr.size; i++ {
		str += fmt.Sprintf("%v", arr.data[i])
		if i != arr.size-1 {
			str += ", "
		}
	}
	str += "]"

	return str
}

// resize : 扩/缩 容
func (arr *ArrayData) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity, newCapacity)
	for i := 0; i < arr.size; i++ {
		newArr[i] = arr.data[i]
	}

	arr.data = newArr
}
