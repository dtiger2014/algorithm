package util

type arrayStack struct {
	data *myArray
}

func ArrayStack(cap int) *arrayStack {
	return &arrayStack{
		data: MyArray(cap),
	}
}

func (this *arrayStack) GetSize() int {
	return this.data.GetSize()
}

func (this *arrayStack) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *arrayStack) GetCapacity() int {
	return this.data.GetCapacity()
}

func (this *arrayStack) Push(e E) {
	this.data.AddLast(e)
}

func (this *arrayStack) Pop() E {
	return this.data.RemoveLast()
}

func (this *arrayStack) Peek() E {
	return this.data.GetLast()
}

func (this *arrayStack) ToString() {
	this.data.ToString()
}