package loopqueue

import (
	"errors"
	"fmt"
	"golang-datastruct/util/common"
)

// LoopQueue : loop queue
type LoopQueue struct {
	data  []interface{}
	tail  int
	front int
	size  int
}

// New : 创建并初始化，并返回LoopQueue
func New() *LoopQueue {
	return &LoopQueue{
		data:  make([]interface{}, common.DefaultCap+1, common.DefaultCap+1),
		tail:  0,
		front: 0,
		size:  0,
	}
}

// GetCapacity : 获取Capacity
func (queue *LoopQueue) GetCapacity() int {
	return cap(queue.data) - 1
}

// GetSize : 获取Size
func (queue *LoopQueue) GetSize() int {
	return queue.size
}

// IsEmpty : 判断是否为空
func (queue *LoopQueue) IsEmpty() bool {
	return queue.tail == queue.front
}

// Enqueue : 入队
func (queue *LoopQueue) Enqueue(e interface{}) {
	if (queue.tail+1)%cap(queue.data) == queue.front {
		queue.resize(queue.GetCapacity() * 2)
	}

	queue.data[queue.tail] = e
	queue.tail = (queue.tail + 1) % cap(queue.data)
	queue.size++
}

// Dequeue : 出队
func (queue *LoopQueue) Dequeue() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Cannot dequeue from an empty queue")
	}

	ret := queue.data[queue.front]
	queue.data[queue.front] = nil
	queue.front = (queue.front + 1) % cap(queue.data)
	queue.size--
	if queue.size == queue.GetCapacity()/4 && queue.GetCapacity()/2 != 0 {
		queue.resize(queue.GetCapacity() / 2)
	}

	return ret, nil
}

// GetFront : 获取队首元素
func (queue *LoopQueue) GetFront() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	return queue.data[queue.front], nil
}

// resize : 扩/缩容
func (queue *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity, newCapacity)

	for i := 0; i < queue.size; i++ {
		newData[i] = queue.data[(i+queue.front)%cap(queue.data)]
	}

	queue.data = newData
	queue.front = 0
	queue.tail = queue.size
}

// String : 格式化输出
func (queue *LoopQueue) String() string {
	str := fmt.Sprintf("Queue: size=%d, capacity = %d\n", queue.size, queue.GetCapacity())
	str += "front ["
	for i := queue.front; i != queue.tail; i = (i + 1) % cap(queue.data) {
		str += fmt.Sprintf("%v", queue.data[i])
		if (i+1)%cap(queue.data) != queue.tail {
			str += ", "
		}
	}
	str += "] tail"

	return str
}
