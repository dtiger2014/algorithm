package unionfind

import "errors"

// UnionFind3 :
type UnionFind3 struct {
	parent []int // parent[i]表示第一个元素所指向的父节点
	sz     []int // sz[i]表示以i为根的集合中元素个数
}

// NewUnionFind3 初始化
func NewUnionFind3(size int) *UnionFind3 {
	uf := &UnionFind3{
		parent: make([]int, 0, size),
		sz:     make([]int, 0, size),
	}

	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < size; i++ {
		uf.parent = append(uf.parent, i)
		uf.sz = append(uf.sz, 1)
	}

	return uf
}

// GetSize 获取 size
func (uf *UnionFind3) GetSize() int {
	return len(uf.parent)
}

// find 查找过程, 查找元素p所对应的集合编号, O(h)复杂度, h为树的高度
func (uf *UnionFind3) find(p int) (int, error) {
	if p < 0 || p > len(uf.parent) {
		return 0, errors.New("p is out of bound")
	}

	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p, nil
}

// IsConnected 查看元素p和元素q是否所属一个集合, O(h)复杂度, h为树的高度
func (uf *UnionFind3) IsConnected(p, q int) bool {
	pID, _ := uf.find(p)
	qID, _ := uf.find(q)

	return pID == qID
}

// UnionElements 合并元素p和元素q所属的集合, O(h)复杂度, h为树的高度
func (uf *UnionFind3) UnionElements(p, q int) {
	pRoot, _ := uf.find(p)
	qRoot, _ := uf.find(q)

	if pRoot == qRoot {
		return
	}
	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if uf.sz[pRoot] < uf.sz[qRoot] {
		uf.parent[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	} else {
		uf.parent[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}
}
