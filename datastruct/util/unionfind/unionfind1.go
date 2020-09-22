package unionfind

import "errors"

// UnionFind1 : 并查集（数组实现）
type UnionFind1 struct {
	id []int
}

// NewUnionFind1 : 初始化UnionFind1
func NewUnionFind1(size int) *UnionFind1 {
	uf := &UnionFind1{
		id: make([]int, 0, size),
	}

	// 每一个id[i]指向自己, 没有合并的元素
	for i := 0; i < size; i++ {
		uf.id = append(uf.id, i)
	}
	return uf
}

// GetSize : 获取size
func (uf *UnionFind1) GetSize() int {
	return len(uf.id)
}

// find : 查找元素p所对应的集合编号, O(1)复杂度
func (uf *UnionFind1) find(p int) (int, error) {
	if p < 0 || p >= len(uf.id) {
		return 0, errors.New("p is out of bound")
	}

	return uf.id[p], nil
}

// IsConnected : 查看元素p和元素q是否所属一个集合,O(1)复杂度
func (uf *UnionFind1) IsConnected(p, q int) bool {
	pID, _ := uf.find(p)
	qID, _ := uf.find(q)
	return pID == qID
}

// UnionElements : 合并元素p和元素q所属的集合, O(n) 复杂度
func (uf *UnionFind1) UnionElements(p, q int) {
	pID, _ := uf.find(p)
	qID, _ := uf.find(q)

	if pID == qID {
		return
	}

	// 合并过程需要遍历一遍所有元素, 将两个元素的所属集合编号合并
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
}