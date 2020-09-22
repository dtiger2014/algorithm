package util

import(
	"fmt"
)

type linkedList struct {
	dummyHead *node
	size int
}

func LinkedList() *linkedList {
	return &linkedList{
		dummyHead: &node{},
		size: 0,
	}
}

func (this *linkedList) GetSize() int {
	return this.size
}

func (this *linkedList) IsEmpty() bool {
	return this.size == 0
}

func (this *linkedList) AddFirst(e E) {
	this.Add(0, e)
}

func (this *linkedList) AddLast(e E) {
	this.Add(this.size, e)
}

func (this *linkedList) Add(index int, e E) {
	if index < 0 || index > this.size {
		panic("Add failed. Illegal index.")
	}

	prev := this.dummyHead;
	for i := 0 ; i < index; i++ {
		prev = prev.next
	}

	prev.next = &node{e, prev.next}
	this.size++
}

func (this *linkedList) Get(index int) E {
	if index < 0 || index >= this.size {
		panic("Get failed. Illegal index.")
	}

	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (this *linkedList) GetFirst() E {
	return this.Get(0)
}

func (this *linkedList) GetLast() E {
	return this.Get(this.size - 1)
}

func (this *linkedList) Set(index int, e E) {
	if index < 0 || index > this.size {
		panic("Set failed. Illegal index.")
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (this *linkedList) Contains(e E) bool {
	for cur := this.dummyHead.next; cur != nil; cur = cur.next {
		if cur.e == e {
			return true
		}
	} 
	return false
}

func (this *linkedList) Remove(index int) E {
	if index < 0 || index >= this.size {
		panic("Remove failed. Index is illegal.")
	}

	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	this.size--

	return retNode.e
}

func (this *linkedList) RemoveFirst() E {
	return this.Remove(0)
}

func (this *linkedList) RemoveLast() E {
	return this.Remove(this.size - 1)
}

func (this *linkedList) RemoveElement(e E) {
	prev := this.dummyHead;
	for prev.next != nil {
		if prev.next.e == e {
			break
		}
		prev = prev.next
	}

	if prev != nil {
		delNode := prev.next
		prev.next = delNode.next;
		delNode.next = nil
		this.size--
	}
}

func (this *linkedList) ToString() {
	fmt.Printf("size=%v [", this.size)
	for cur := this.dummyHead.next; cur != nil; cur = cur.next {
		fmt.Printf("%v => ", cur.e)
	}
	fmt.Println(" nil]")
}
