package util

import (
	"fmt"
)

type loopQueue struct {
	data []E
	front, tail, size int
}

func LoopQueue(cap int) *loopQueue {
	return &loopQueue{
		data: make([]E, cap+1),
		front: 0,
		tail: 0,
		size: 0,
	}
}

func (this *loopQueue) GetSize() int {
	return this.size
}

func (this *loopQueue) IsEmpty() bool {
	return this.front == this.tail
}

func (this *loopQueue) GetCapacity() int {
	return cap(this.data) - 1
}

func (this *loopQueue) Enqueue(e E) {
	if (this.tail + 1) % cap(this.data) == this.front {
		this.resize(this.GetCapacity() * 2)
	}

	this.data[this.tail] = e
	this.tail = (this.tail + 1) % cap(this.data)
	this.size++;
}

func (this *loopQueue) Dequeue() E {
	if this.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	ret := this.data[this.front]
	this.data[this.front] = nil
	this.front = (this.front + 1) % cap(this.data)
	this.size--
	if this.size == this.GetCapacity() / 4 && this.GetCapacity() / 2 != 0 {
		this.resize(this.GetCapacity() / 2)
	}

	return ret
}

func (this *loopQueue) GetFront() E {
	if this.IsEmpty() {
		panic("LoopQueue is empty.")
	}
	return this.data[this.front]
}

func (this *loopQueue) resize(newCapacity int) {
	oldData := this.data
	this.data = make([]E, newCapacity + 1)
	for i := 0; i < this.size; i++ {
		this.data[i] = oldData[(i + this.front) % cap(oldData)]
	}

	this.front = 0
	this.tail = this.size
}

func (this *loopQueue) ToString() {
	fmt.Printf("LoopQueue: size = %d , capacity = %d\n", this.size, this.GetCapacity())
	fmt.Printf("front [")
	for i := this.front; i != this.tail; i = (i + 1) % cap(this.data) {
		fmt.Print(this.data[i])
		if (i + 1) % cap(this.data) != this.tail {
			fmt.Print(", ")
		}
	}
	fmt.Print("] tail\n")
}

