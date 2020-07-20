package sort

func isSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.CompareTo(i, i-1) < 0 {
			return false
		}
	}
	return true
}
