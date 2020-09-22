package avlset

import (
	"golang-datastruct/util/avltree"
)

type AVLSet struct {
	avl *avltree.AVLTree
}

func New() *AVLSet {
	return &AVLSet{
		avl: avltree.New(),
	}
}

func (set *AVLSet) GetSize() int {
	return set.avl.GetSize()
}

func (set *AVLSet) IsEmpty() bool {
	return set.avl.IsEmpty()
}

func (set *AVLSet) Add(e string) {
	set.avl.Add(e, 0)
}

func (set *AVLSet) Contains(e string) bool {
	return set.avl.Contains(e)
}

func (set *AVLSet) Remove(e string) {
	set.avl.Remove(e)
}
