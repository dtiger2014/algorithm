package maxheap

import (
	"errors"
	"golang-datastruct/util/array"
)

type MaxHeap struct {
	data *array.ArrayData
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: array.New(),
	}
}

func Heapify(arr []int) {

}

func (heap *MaxHeap) GetSize() int {
	return heap.data.GetSize()
}

func (heap *MaxHeap) IsEmpty() bool {
	return heap.data.IsEmpty()
}

func (heap *MaxHeap) parent(index int) (int, error) {
	if index == 0 {
		return 0, errors.New("index-0 doesn't have parent")
	}

	return (index - 1) / 2, nil
}

func (heap *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}
func (heap *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

func (heap *MaxHeap) Add(e interface{}) {
	heap.data.AddLast(e)
	heap.siftUp(heap.data.GetSize() - 1)
}

func (heap *MaxHeap) siftUp(k int) {

	// TODO:
	// while(k > 0 && data.get(parent(k)).compareTo(data.get(k)) < 0 ){
	// 	data.swap(k, parent(k));
	// 	k = parent(k);
	// }

	// parentKey, _ := heap.parent(k)
	// parent, _ := heap.data.Get(parentKey)
	// cur, _ := heap.data.Get(k)
	// for k > 0 && parent < cur {

	// 	heap.data.Swap(k, parent)
	// }
}

func (heap *MaxHeap) FindMax() (interface{}, error) {
	if heap.data.GetSize() == 0 {
		return nil, errors.New("Can not findMax when heap is empty")
	}

	return heap.data.Get(0)
}

// ExtractMax : 取出堆中最大元素
func (heap *MaxHeap) ExtractMax() interface{} {
	ret, _ := heap.FindMax()
	heap.data.Swap(0, heap.data.GetSize()-1)
	heap.data.RemoveLast()
	heap.siftDown(0)

	return ret
}

func (heap *MaxHeap) siftDown(k int) {
	// while(leftChild(k) < data.getSize()){
	// 	int j = leftChild(k); // 在此轮循环中,data[k]和data[j]交换位置
	// 	if( j + 1 < data.getSize() &&
	// 			data.get(j + 1).compareTo(data.get(j)) > 0 )
	// 		j ++;
	// 	// data[j] 是 leftChild 和 rightChild 中的最大值

	// 	if(data.get(k).compareTo(data.get(j)) >= 0 )
	// 		break;

	// 	data.swap(k, j);
	// 	k = j;
	// }
}

func (heap *MaxHeap) Replace(e interface{}) interface{} {
	ret, _ := heap.FindMax()
	heap.data.Set(0, e)
	heap.siftDown(0)
	return ret
}
