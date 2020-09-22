package linkedlist

import (
	"errors"
	"fmt"
)

// Node : 链表元素
type Node struct {
	value interface{}
	next  *Node
}

// newNode : 创建初始化，并返回 Node
func newNode(e interface{}, node *Node) *Node {
	return &Node{
		value: e,
		next:  node,
	}
}

// LinkedList : 链表
type LinkedList struct {
	dummyHead *Node
	size      int
}

// New : 创建初始化，并返回 LinkedList
func New() *LinkedList {
	return &LinkedList{
		dummyHead: newNode(nil, nil),
		size:      0,
	}
}

// GetSize : 获取Size
func (l *LinkedList) GetSize() int {
	return l.size
}

// IsEmpty : 判断是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Add : 添加元素
func (l *LinkedList) Add(index int, e interface{}) error {
	if index < 0 || index > l.size {
		return errors.New("Add failed. Illegal index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = newNode(e, prev.next)
	l.size++

	return nil
}

// AddFirst : 添加元素 First
func (l *LinkedList) AddFirst(e interface{}) error {
	return l.Add(0, e)
}

// AddLast : 添加元素 Last
func (l *LinkedList) AddLast(e interface{}) error {
	return l.Add(l.size, e)
}

// Get : 获取元素
func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("Get failed. Illegal index")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur, nil
}

// GetFirst : 获取元素 First
func (l *LinkedList) GetFirst() (interface{}, error) {
	return l.Get(0)
}

// GetLast : 获取元素 Last
func (l *LinkedList) GetLast() (interface{}, error) {
	return l.Get(l.size - 1)
}

// Set : 设置元素
func (l *LinkedList) Set(index int, e interface{}) error {
	if index < 0 || index >= l.size {
		return errors.New("Update failed. Illegal index")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	cur.value = e

	return nil
}

// Contains ：判断是否包含目标元素
func (l *LinkedList) Contains(e interface{}) bool {
	for cur := l.dummyHead.next; cur.next != nil; cur = cur.next {
		if cur.value == e {
			return true
		}
	}
	return false
}

// Remove : 删除目标索引元素
func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("Remove failed. Index is illegal")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--

	return retNode.value, nil
}

// RemoveFirst : 删除元素 First
func (l *LinkedList) RemoveFirst() (interface{}, error) {
	return l.Remove(0)
}

// RemoveLast : 删除元素 Last
func (l *LinkedList) RemoveLast() (interface{}, error) {
	return l.Remove(l.size - 1)
}

// RemoveElement : 删除目标元素
func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead
	for ; prev.next != nil; prev = prev.next {
		if prev.next.value == e {
			break
		}
	}

	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		l.size--
	}
}

// String : 格式化输出
func (l *LinkedList) String() string {
	str := ""

	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		str += fmt.Sprintf("%v -> ", cur.value)
	}

	str += "nil"

	return str
}
