package unionfind

import "errors"

// UnionFind6 :
type UnionFind6 struct {
	// rank[i]表示以i为根的集合所表示的树的层数
	// 在后续的代码中, 我们并不会维护rank的语意, 也就是rank的值在路径压缩的过程中, 有可能不在是树的层数值
	// 这也是我们的rank不叫height或者depth的原因, 他只是作为比较的一个标准
	rank []int

	parent []int // parent[i]表示第一个元素所指向的父节点

}

// NewUnionFind6 初始化
func NewUnionFind6(size int) *UnionFind6 {
	uf := &UnionFind6{
		parent: make([]int, 0, size),
		rank:   make([]int, 0, size),
	}

	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < size; i++ {
		uf.parent = append(uf.parent, i)
		uf.rank = append(uf.rank, 1)
	}

	return uf
}

// GetSize 获取 size
func (uf *UnionFind6) GetSize() int {
	return len(uf.parent)
}

// find 查找过程, 查找元素p所对应的集合编号, O(h)复杂度, h为树的高度
func (uf *UnionFind6) find(p int) (int, error) {
	if p < 0 || p > len(uf.parent) {
		return 0, errors.New("p is out of bound")
	}

	// path compression 2, 递归算法
	for p != uf.parent[p] {
		uf.parent[p], _ = uf.find(uf.parent[p])
	}
	return uf.parent[p], nil
}

// IsConnected 查看元素p和元素q是否所属一个集合, O(h)复杂度, h为树的高度
func (uf *UnionFind6) IsConnected(p, q int) bool {
	pID, _ := uf.find(p)
	qID, _ := uf.find(q)

	return pID == qID
}

// UnionElements 合并元素p和元素q所属的集合, O(h)复杂度, h为树的高度
func (uf *UnionFind6) UnionElements(p, q int) {
	pRoot, _ := uf.find(p)
	qRoot, _ := uf.find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = qRoot
	} else if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else {
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot]++
	}
}
