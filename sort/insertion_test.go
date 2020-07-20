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

func Test_InsertionSort(t *testing.T) {
	data := ints
	a := SortInt(data)
	InsertionSort(a)
	if !isSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", a)
	}
}

func Benchmark_InsertionSort(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		n := 10000
		randArr := utils.GenerateRandomArray(n, 0, n)
		a := SortInt(randArr)

		b.StartTimer()
		InsertionSort(a)
		b.StopTimer()

		if !isSorted(a) {
			b.Errorf("Benchmark_InsertionSort is not sorted! n=%d", n)
		}
	}
}

func Test_InsertionSort2(t *testing.T) {
	data := ints
	a := SortInt(data)
	InsertionSort2(a)
	if !isSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", a)
	}
}

func Test_InsertionSortInts(t *testing.T) {
	n := 10000
	randArr := utils.GenerateRandomArray(n, 0, n)

	InsertionSortInts(randArr)

	if !isSorted(SortInt(randArr)) {
		t.Errorf("Test_InsertionSortInts is not sorted! n=%d", n)
	}
}

func Benchmark_InsertionSortInts(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		n := 10000

		randArr := utils.GenerateRandomArray(n, 0, n)

		b.StartTimer()
		InsertionSortInts(randArr)
		b.StopTimer()

		if !isSorted(SortInt(randArr)) {
			b.Errorf("Benchmark_InsertionSortInts is not sorted! n=%d", n)
		}
	}
}
