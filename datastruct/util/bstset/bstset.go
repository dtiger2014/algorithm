package bstset

import (
	"golang-datastruct/util/bst"
)

type BSTSet struct {
	data *bst.BST
}

func New() *BSTSet {
	return &BSTSet{
		data: bst.New(),
	}
}

func (set *BSTSet) GetSize() int {
	return set.data.GetSize()
}

func (set *BSTSet) IsEmpty() bool {
	return set.data.IsEmpty()
}

func (set *BSTSet) Add(e int) {
	set.data.Add(e)
}

func (set *BSTSet) Contains(e int) bool {
	return set.data.Contains(e)
}

func (set *BSTSet) Remove(e int) {
	set.data.Remove(e)
}