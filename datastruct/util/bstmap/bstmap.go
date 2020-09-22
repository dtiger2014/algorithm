package bstmap

import (
	"errors"
	"fmt"
)

type node struct {
	key         string
	value       int
	left, right *node
}

type BSTMap struct {
	root *node
	size int
}

func New() *BSTMap {
	return &BSTMap{
		root: &node{},
		size: 0,
	}
}

func (m *BSTMap) GetSize() int {
	return m.size
}

func (m *BSTMap) IsEmpty() bool {
	return m.size == 0
}

func (m *BSTMap) Add(key string, value int) {
	m.root = m.add(m.root, key, value)
}

func (m *BSTMap) add(n *node, key string, value int) *node {
	if n == nil {
		m.size++
		return &node{
			key:   key,
			value: value,
		}
	}

	if key < n.key {
		n.left = m.add(n.left, key, value)
	} else if key > n.key {
		n.right = m.add(n.right, key, value)
	} else {
		n.value = value
	}

	return n
}

func (m *BSTMap) GetNode(n *node, k string) *node {
	if n == nil {
		return nil
	}

	if k == n.key {
		return n
	} else if k < n.key {
		return m.GetNode(n.left, k)
	} else {
		return m.GetNode(n.right, k)
	}
}

func (m *BSTMap) Contains(k string) bool {
	return m.GetNode(m.root, k) != nil
}

func (m *BSTMap) Get(k string) (int, error) {
	node := m.GetNode(m.root, k)
	if node == nil {
		errstr := fmt.Sprintf("Key: %s Not Found", k)
		return 0, errors.New(errstr)
	} else {
		return node.value, nil
	}
}

func (m *BSTMap) Set(k string, v int) error {
	node := m.GetNode(m.root, k)
	if node == nil {
		errstr := fmt.Sprintf("Key:%s doesn't exists", k)
		return errors.New(errstr)
	}

	node.value = v
	return nil
}

func (m *BSTMap) minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return m.minimum(n.left)
}

func (m *BSTMap) removeMin(n *node) *node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		m.size--
		return rightNode
	}

	n.left = m.removeMin(n.left)
	return n
}

func (m *BSTMap) Remove(k string) (int, error) {
	node := m.GetNode(m.root, k)
	if node != nil {
		m.root = m.remove(m.root, k)
		return node.value, nil
	}
	return 0, errors.New("Not Found")
}

func (m *BSTMap) remove(n *node, k string) *node {
	if n == nil {
		return nil
	}

	if k < n.key {
		n.left = m.remove(n.left, k)
		return n
	} else if k > n.key {
		n.right = m.remove(n.right, k)
		return n
	} else {

		if n.left == nil {
			rightNode := n.right
			n.right = nil
			m.size--
			return rightNode
		}

		if n.right == nil {
			leftNode := n.left
			n.left = nil
			m.size--
			return leftNode
		}

		successor := m.minimum(n.right)
		successor.right = m.removeMin(n.right)
		successor.left = n.left

		n.left = nil
		n.right = nil

		return successor
	}
}
