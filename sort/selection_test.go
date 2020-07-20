package sort

import (
	"algorithm/utils"
	"testing"
)

var (
	ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
)

type SortInt []int

func (p SortInt) Len() int               { return len(p) }
func (p SortInt) CompareTo(i, j int) int { return p[i] - p[j] }
func (p SortInt) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }

func Test_SelectionSort(t *testing.T) {
	data := ints
	a := SortInt(data)
	SelectionSort(a)
	if !isSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", a)
	}
}

func Test_SelectionSort2(t *testing.T) {
	data := ints
	a := SortInt(data)
	SelectionSort2(a)
	if !isSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", a)
	}
}

func Benchmark_SelectionSort(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		n := 10000
		randArr := utils.GenerateRandomArray(n, 0, n)
		a := SortInt(randArr)

		b.StartTimer()
		SelectionSort(a)
		b.StopTimer()

		if !isSorted(a) {
			b.Errorf("Benchmark_SelectionSort is not sorted! n=%d", n)
		}
	}
}

func Benchmark_SelectionSort2(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		n := 10000
		randArr := utils.GenerateRandomArray(n, 0, n)
		a := SortInt(randArr)

		b.StartTimer()
		SelectionSort2(a)
		b.StopTimer()

		if !isSorted(a) {
			b.Errorf("Benchmark_SelectionSort2 is not sorted! n=%d", n)
		}
	}
}
