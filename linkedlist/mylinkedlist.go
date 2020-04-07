package linkedlist

import "fmt"

type MyNode struct {
	val  int
	next *MyNode
}
type MyLinkedList struct {
	dummy *MyNode
	size  int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		dummy: &MyNode{},
		size:  0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.size == 0 || index >= this.size {
		return -1
	}

	dumm := this.dummy
	for i := 0; i <= index; i++ {
		dumm = dumm.next
	}
	return dumm.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		return
	}
	newNode := &MyNode{
		val:  val,
		next: nil,
	}

	dumm := this.dummy
	for i := 0; i < index; i++ {
		dumm = dumm.next
	}

	newNode.next = dumm.next
	dumm.next = newNode

	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	dumm := this.dummy
	for i := 0; i < index; i++ {
		dumm = dumm.next
	}

	delNode := dumm.next
	dumm.next = dumm.next.next
	delNode.next = nil
	this.size--
}

func (this *MyLinkedList) String() string {
	res := ""
	node := this.dummy.next
	for i := 0; i < this.size; i++ {
		res += fmt.Sprintf("%v-> ", node.val)
		node = node.next
	}
	return res
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
