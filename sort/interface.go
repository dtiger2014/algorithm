package sort

// Interface 定义排序 接口
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
