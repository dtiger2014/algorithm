package util

import(
	"fmt"
)

type linkedListQueue struct {
	head, tail *node
	size int
}

func LinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *linkedListQueue) GetSize() int {
	return this.size
}

func (this *linkedListQueue) IsEmpty() bool {
	return this.size == 0
}


func (this *linkedListQueue) Enqueue(e E) {
	if this.tail == nil {
		this.tail = &node{e, nil}
		this.head = this.tail
	} else {
		this.tail.next = &node{e, nil}
		this.tail = this.tail.next
	}
	this.size++
}

func (this *linkedListQueue) Dequeue() E {
	if this.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	retNode := this.head
	this.head = this.head.next
	retNode.next = nil
	if this.head == nil {
		this.tail = nil
	}
	this.size--
	return retNode.e
}

func (this *linkedListQueue) GetFront() E {
	if this.IsEmpty() {
		panic("Queue is empty.")
	}
	return this.head.e
}

func (this *linkedListQueue) ToString() {
	fmt.Print("LinkedListQueue: head [")
	for cur := this.head; cur != nil; cur = cur.next {
		fmt.Printf("%v=>", cur.e)
	}
	fmt.Println("nil] tail")
}