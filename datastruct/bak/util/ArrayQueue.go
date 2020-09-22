package util

type arrayQueue struct {
	data *myArray
}

func ArrayQueue(cap int) *arrayQueue {
	return &arrayQueue{
		data: MyArray(cap),
	}
}

func (this *arrayQueue) GetSize() int {
	return this.data.GetSize()
}

func (this *arrayQueue) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *arrayQueue) GetCapacity() int {
	return this.data.GetCapacity()
}

func (this *arrayQueue) Enqueue(e E) {
	this.data.AddLast(e)
}

func (this *arrayQueue) Dequeue() E {
	return this.data.RemoveFirst()
}

func (this *arrayQueue) GetFront() E {
	return this.data.GetFirst()
}

func (this *arrayQueue) ToString() {
	this.data.ToString()
}