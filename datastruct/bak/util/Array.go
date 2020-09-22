package util

/*
result:
	RemoveLast() : O(1)
	RemoveFirst() : O(n)
	AddLast() : O(1)
	AddFirst() : O(n)

	Last 和 First 相差 700倍
*/

import(
	"fmt"
)

type myArray struct {
	data []E
}

func MyArray(cap int) *myArray {
	return &myArray{
		data: make([]E, 0, cap),
	}
}

func (this *myArray) GetCapacity() int{
	return cap(this.data)
}

func (this *myArray) GetSize() int {
	return len(this.data)
}

func (this *myArray) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *myArray) AddFirst(e E) {
	this.Add(0, e)
}

func (this *myArray) AddLast(e E) {
	this.Add(len(this.data), e)
}

func (this *myArray) Add(index int, e E) {
	if index < 0 || index > len(this.data) {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if len(this.data) == cap(this.data) {
		// panic("Add failed. Array is full")
		this.resize(2 * cap(this.data))
	}
	
	// 保存 slice 中 index 后面的元素
	slice := append([]E{}, this.data[index:]...)
	this.data = append(append(this.data[:index], e), slice...)
}

func (this *myArray) Set(index int, e E) {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	this.data[index] = e
}

func (this *myArray) Get(index int) E {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

func (this *myArray) GetLast() E {
	return this.Get(len(this.data) - 1)
}

func (this *myArray) GetFirst() E {
	return this.Get(0)
}

func (this *myArray) Contains(e E) bool {
	for _, v := range this.data {
		if v == e {
			return true
		}
	}
	return false
}

func (this *myArray) Find(e E) int {
	for i, v := range this.data {
		if v == e {
			return i
		}
	}
	return -1
}

func (this *myArray) Remove(index int) E {
	if index < 0 || index >= len(this.data) {
		panic("Remove failed. Index is illegal")
	}

	ret := this.data[index]
	this.data = append(this.data[:index], this.data[index+1:]...)

	if (len(this.data) == cap(this.data) / 4 && len(this.data) / 2 != 0) {
		this.resize(cap(this.data) / 2)
	}

	return ret
}

func (this *myArray) RemoveFirst() E {
	return this.Remove(0)
}

func (this *myArray) RemoveLast() E {
	return this.Remove(len(this.data) - 1)
}

func (this *myArray) RemoveElement(e E) {
	index := this.Find(e)
	if index != -1 {
		this.Remove(index)
	}
}

func (this *myArray) resize(newCapacity int) {
	slice := this.data
	this.data = make([]E, 0, newCapacity)
	this.data = append(this.data, slice...)
}

func (this *myArray) ToString() {
	fmt.Printf("Array: size = %v , capacity = %v\n", len(this.data), cap(this.data))
	fmt.Println(this.data)
}