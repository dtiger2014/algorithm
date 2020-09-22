package avlmap

import (
	"golang-datastruct/util/avltree"
)

type AVLMap struct {
	avl *avltree.AVLTree
}

func New() *AVLMap {
	return &AVLMap{
		avl: avltree.New(),
	}
}

func (m *AVLMap) GetSize() int {
	return m.avl.GetSize()
}

func (m *AVLMap) IsEmpty() bool {
	return m.avl.IsEmpty()
}

func (m *AVLMap) Add(e string) {
	m.avl.Add(e, 0)
}

func (m *AVLMap) Contains(e string) bool {
	return m.avl.Contains(e)
}

func (m *AVLMap) Remove(e string) {
	m.avl.Remove(e)
}
