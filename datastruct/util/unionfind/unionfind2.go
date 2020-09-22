package unionfind

import "errors"

// UnionFind2 : 我们的第二版Union-Find, 使用一个数组构建一棵指向父节点的树
// parent[i]表示第一个元素所指向的父节点
type UnionFind2 struct {
	parent []int
}

// NewUnionFind2 初始化
func NewUnionFind2(size int) *UnionFind2 {
	uf := &UnionFind2{
		parent: make([]int, 0, size),
	}

	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < size; i++ {
		uf.parent = append(uf.parent, i)
	}

	return uf
}

// GetSize 获取 size
func (uf *UnionFind2) GetSize() int {
	return len(uf.parent)
}

// find 查找过程, 查找元素p所对应的集合编号, O(h)复杂度, h为树的高度
func (uf *UnionFind2) find(p int) (int, error) {
	if p < 0 || p > len(uf.parent) {
		return 0, errors.New("p is out of bound")
	}

	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p, nil
}

// IsConnected 查看元素p和元素q是否所属一个集合, O(h)复杂度, h为树的高度
func (uf *UnionFind2) IsConnected(p, q int) bool {
	pID, _ := uf.find(p)
	qID, _ := uf.find(q)

	return pID == qID
}

// UnionElements 合并元素p和元素q所属的集合, O(h)复杂度, h为树的高度
func (uf *UnionFind2) UnionElements(p, q int) {
	pRoot, _ := uf.find(p)
	qRoot, _ := uf.find(q)

	if pRoot == qRoot {
		return
	}

	uf.parent[pRoot] = qRoot
}
