package util

import(
	"fmt"
)

type E interface{}

type node struct {
	e E
	next *node
}

func (this *node) toString() {
	fmt.Println(this.e)
}

type Operation interface{
	GetCapacity() int
	GetSize() int
	IsEmpty() bool
	ToString()

	AddFirst(e interface{})
	AddLast(e interface{})
	Add(index int, e interface{})
	Set(index int, e interface{})
	Get(index int) interface{}
	Contains(e interface{} ) bool
	Find(e interface{}) int
	Remove(index int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	RemoveElement(e interface{})

	Push()
    Pop()
    Peek()
}
