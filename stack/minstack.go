package stack

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 || x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.data[len(this.data)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}

	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
