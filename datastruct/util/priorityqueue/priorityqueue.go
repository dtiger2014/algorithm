package priorityqueue

import (
	"golang-datastruct/util/maxheap"
)

type PriorityQueue struct {
	data *maxheap.MaxHeap
}

func (q *PriorityQueue) GetSize() int {
	return q.data.GetSize()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *PriorityQueue) GetFront() (interface{}, error) {
	return q.data.FindMax()
}

func (q *PriorityQueue) Enqueue(e interface{}) {
	q.data.Add(e)
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.data.ExtractMax()
}
