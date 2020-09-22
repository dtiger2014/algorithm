package linkedlistmap

import (
	"errors"
	"fmt"
)

type Node struct {
	Key   string
	Value int
	Next  *Node
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) String() string {
	return fmt.Sprintf("%s : %d", node.Key, node.Value)
}

type LinkedListMap struct {
	dummyHead *Node
	size      int
}

func New() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: NewNode(),
		size:      0,
	}
}

func (m *LinkedListMap) GetSize() int {
	return m.size
}

func (m *LinkedListMap) IsEmpty() bool {
	return m.size == 0
}

func (m *LinkedListMap) GetNode(key string) *Node {
	cur := m.dummyHead.Next

	for cur != nil {
		if cur.Key == key {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (m *LinkedListMap) Contains(key string) bool {
	return m.GetNode(key) != nil
}

func (m *LinkedListMap) Get(key string) (int, error) {
	node := m.GetNode(key)
	if node == nil {
		return 0, errors.New("Not Found")
	} else {
		return node.Value, nil
	}
}

func (m *LinkedListMap) Add(key string, value int) {
	node := m.GetNode(key)
	if node == nil {
		m.dummyHead.Next = &Node{
			Key:   key,
			Value: value,
			Next:  m.dummyHead.Next,
		}
		m.size++
	} else {
		node.Value = value
	}
}

func (m *LinkedListMap) Set(key string, value int) error {
	node := m.GetNode(key)

	if node == nil {
		errstr := fmt.Sprintf("Key: %s doesn't exist", key)
		return errors.New(errstr)
	}

	node.Value = value

	return nil
}

func (m *LinkedListMap) Remove(key string) (int, error) {
	prev := m.dummyHead
	for prev.Next != nil {
		if prev.Next.Key == key {
			break
		}
		prev = prev.Next
	}

	if prev.Next != nil {
		delNode := prev.Next
		prev.Next = delNode.Next
		delNode.Next = nil
		m.size--

		return delNode.Value, nil
	}

	return 0, errors.New("Remove Element Not Found")
}
