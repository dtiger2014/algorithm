package sort

// Interface 定义排序 接口
type Interface interface {
	Len() int
	CompareTo(i, j int) int
	Swap(i, j int)
}
