package util

import(
	"fmt"
)

type linkedListStack struct {
	data *linkedList
}

func LinkedListStack() *linkedListStack {
	return &linkedListStack{
		data : LinkedList(),
	}
}

func (this *linkedListStack) GetSize() int {
	return this.data.GetSize()
}

func (this *linkedListStack) IsEmpty() bool {
	return this.data.IsEmpty()
}


func (this *linkedListStack) Push(e E) {
	this.data.AddFirst(e)
}

func (this *linkedListStack) Pop() E {
	return this.data.RemoveFirst()
}

func (this *linkedListStack) Peek() E {
	return this.data.GetFirst()
}

func (this *linkedListStack) ToString() {
	fmt.Print("Stack: top")
	this.data.ToString()
}