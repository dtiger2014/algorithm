package heap

import "errors"

// MaxHeap 最大堆
type MaxHeap struct {
	data  []int
	count int
}

// NewMaxHeap 创建最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data:  make([]int, 1),
		count: 0,
	}
}

// NewMaxHeapByArray 根据传入的数组创建堆
func NewMaxHeapByArray(arr []int) *MaxHeap {
	heap := &MaxHeap{
		data:  make([]int, len(arr)+1),
		count: len(arr),
	}

	for i, v := range arr {
		heap.data[i+1] = v
	}

	for i := heap.count / 2; i >= 1; i-- {
		heap.shiftDown(i)
	}

	return heap
}

// Size 元素个数
func (h *MaxHeap) Size() int {
	return h.count
}

// IsEmpty 是否为空？
func (h *MaxHeap) IsEmpty() bool {
	return h.count == 0
}

// Insert 入队
func (h *MaxHeap) Insert(item int) {
	h.data = append(h.data, item)
	h.shiftUp(h.count + 1)
	h.count++
}

func (h *MaxHeap) shiftUp(k int) {
	for {
		if !(k > 1 && h.data[k/2] < h.data[k]) {
			break
		}
		h.data[k/2], h.data[k] = h.data[k], h.data[k/2]
		k /= 2
	}
}

// ExtractMax 推出最大元素
func (h *MaxHeap) ExtractMax() (int, error) {
	if h.count <= 0 {
		return 0, errors.New("MaxHeap is empty")
	}
	ret := h.data[1]

	h.data[1], h.data[h.count] = h.data[h.count], h.data[1]
	h.data = h.data[:h.count]
	h.count--
	h.shiftDown(1)

	return ret, nil
}

func (h *MaxHeap) shiftDown(k int) {
	for {
		if !(k*2 <= h.count) {
			break
		}

		j := 2 * k
		if j+1 <= h.count && h.data[j+1] > h.data[j] {
			j++
		}

		if h.data[k] >= h.data[j] {
			break
		}
		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}

// GetMax 获取最大元素
func (h *MaxHeap) GetMax() (int, error) {
	if h.count <= 0 {
		return 0, errors.New("MaxHeap is empty")
	}

	return h.data[1], nil
}
