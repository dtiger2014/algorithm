package sort

import (
	"algorithm/utils"
	"testing"
)

// var (
// 	ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
// )

// type SortInt []int

// func (p SortInt) Len() int           { return len(p) }
// func (p SortInt) Less(i, j int) bool { return p[i] < p[j] }
// func (p SortInt) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestInsertionSort(t *testing.T) {
	data := ints
	a := SortInt(data)
	InsertionSort(a)
	if !isSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", a)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		n := 10000
		randArr := utils.GenerateRandomArray(n, 0, n)
		a := SortInt(randArr)

		b.StartTimer()
		SelectionSort(a)
		b.StopTimer()

		if !isSorted(a) {
			b.Errorf("BenchmarkInsertionSort is not sorted! n=%d", n)
		}
	}
}
