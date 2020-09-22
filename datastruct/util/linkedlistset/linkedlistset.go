package linkedlistset

import (
	"golang-datastruct/util/linkedlist"
)

type LinkedListSet struct {
	data *linkedlist.LinkedList
}

func New() *LinkedListSet {
	return &LinkedListSet{
		data: linkedlist.New(),
	}
}

func (set *LinkedListSet) GetSize() int {
	return set.data.GetSize()
}

func (set *LinkedListSet) IsEmpty() bool {
	return set.data.IsEmpty()
}

func (set *LinkedListSet) Add(e interface{}) {
	if !set.data.Contains(e) {
		set.data.AddFirst(e)
	}
}

func (set *LinkedListSet) Contains(e interface{}) bool {
	return set.data.Contains(e)
}

func (set *LinkedListSet) Remove(e interface{}) {
	set.data.RemoveElement(e)
}
