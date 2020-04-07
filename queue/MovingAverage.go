package queue

type MovingAverage struct {
	data  []int
	front int
	real  int
	size  int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data:  make([]int, size, size),
		front: 0,
		real:  -1,
		size:  0,
		sum:   0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.size == cap(this.data) {
		this.size--
		this.sum -= this.data[this.front]
		this.front = (this.front + 1) % cap(this.data)
	}
	this.real = (this.real + 1) % cap(this.data)
	this.data[this.real] = val
	this.sum += val
	this.size++

	return float64(this.sum) / float64(this.size)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
