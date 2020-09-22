package segmenttree

type SegmentTree struct {
	tree []int
	data []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	d := &SegmentTree{}
	d.data = make([]int, 0, len(arr))
	d.data = append(d.data, arr...)
	
	d.tree = make([]int, 0, 4 * len(arr))

	// TODO: 

	return d
}