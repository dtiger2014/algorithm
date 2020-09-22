package arrayqueue

import (
	"fmt"
	"golang-datastruct/util/array"
)

// ArrayQueue : array queue
type ArrayQueue struct {
	data *array.ArrayData
}

// New : 创建并初始化，并返回ArrayQueue
func New() *ArrayQueue {
	return &ArrayQueue{
		data: array.New(),
	}
}

// GetCapacity : 获取Capacity
func (queue *ArrayQueue) GetCapacity() int {
	return queue.data.GetCapacity()
}

// GetSize : 获取Size
func (queue *ArrayQueue) GetSize() int {
	return queue.data.GetSize()
}

// IsEmpty : 判断是否为空
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.data.IsEmpty()
}

// Enqueue : 入队
func (queue *ArrayQueue) Enqueue(e interface{}) {
	queue.data.AddLast(e)
}

// Dequeue : 出队
func (queue *ArrayQueue) Dequeue() (interface{}, error) {
	return queue.data.RemoveFirst()
}

// GetFront : 获取队首元素
func (queue *ArrayQueue) GetFront() (interface{}, error) {
	return queue.data.GetFirst()
}

// String : 格式化输出
func (queue *ArrayQueue) String() string {
	str := fmt.Sprintf("Queue:\n")
	str += "front ["
	for i := 0; i < queue.data.GetSize(); i++ {
		v, _ := queue.data.Get(i)
		str += fmt.Sprintf("%v", v)
		if i != queue.data.GetSize()-1 {
			str += ", "
		}
	}
	str += "] tail"

	return str
}
