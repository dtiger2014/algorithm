package data_struct

import(
	"fmt"
	"errors"
	// "math"
)

type MaxHeap struct{
	data 	[]int
	count 	int 
}

func NewMaxHeap(capacity int) *MaxHeap{
	return &MaxHeap{
		data:	make([]int, capacity+1),
		count: 	0,
	}
}

func MaxHeapify(arr []int, n int) *MaxHeap {
	mh := NewMaxHeap(n + 1)

	for i := 0; i < n; i++ {
		mh.data[i+1] = arr[i]
	}
	mh.count = n

	for i := mh.count/2; i >= 1; i-- {
		mh.shiftDown(i)
	}

	return mh
}

func (this *MaxHeap) Size() int {
	return this.count
}

func (this *MaxHeap) IsEmpty() bool {
	return this.count == 0
}

func (this *MaxHeap) Insert(n int) {
	if this.count + 1 > len(this.data) {
		fmt.Printf("MaxHeap is full, size=%v\n", len(this.data))
		return
	}
	this.data[this.count + 1] = n
	this.count++
	this.shiftUp(this.count)
}

func (this *MaxHeap) ExtractMax() (int, error) {
	if this.count <= 0 {
		return 0, errors.New("MaxHeap is Empty")
	}
	ret := this.data[1]
	this.data[1], this.data[this.count] = this.data[this.count], this.data[1]
	this.count--
	
	this.shiftDown(1)

	return ret, nil
}

func (this *MaxHeap) GetMax() (int, error) {
	if this.count <= 0 {
		return 0, errors.New("Max heap is empty")
	}

	return this.data[1], nil
}

func (this *MaxHeap) Print() {
	fmt.Printf("The max heap size is: %v\n", this.Size())
	fmt.Printf("Data in the max heap: ")
	for i := 1; i <= this.Size(); i++ {
		fmt.Printf("%v ", this.data[i])
	}
	fmt.Println()
}



func (this *MaxHeap) shiftUp(k int) {
	for {
		if !(k > 1 && this.data[k / 2] < this.data[k]) {
			break;
		}
		this.data[k/2], this.data[k] = this.data[k], this.data[k/2]
		k = k / 2
	}
}

func (this *MaxHeap) shiftDown(k int) {
	for {
		if !( 2 * k <= this.count) {
			break
		}

		j := 2 * k // 在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= this.count && this.data[j+1] > this.data[j] {
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if this.data[k] >= this.data[j] {
			break
		}

		this.data[k], this.data[j] = this.data[j], this.data[k]
		k = j
	}
}