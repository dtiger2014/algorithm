package util

import(
	// "fmt"
)

type bstnode struct {
	Left, Right *bstnode
	E int	
}

type bst struct {
	root *bstnode
	size int
}

func addBstNode(e int) *bstnode {
	return &bstnode{
		Left: nil,
		Right: nil,
		E: e,
	}
}


func BST() *bst {
	return &bst{
		root: nil,
		size: 0,
	}
}

func (this *bst) GetSize() int {
	return this.size
}

func (this *bst) IsEmpty() bool {
	return this.size == 0
}

func (this *bst) Add(e int) {
	if this.root == nil {
		this.root = addBstNode(e)
		this.size++
	} else {
		this.add(this.root, e)
	}
}
func (this *bst) add(root *bstnode, e int) {
	if root.E == e {
		return
	} else if e < root.E && root.Left == nil {
		root.Left = addBstNode(e)
		this.size++
		return
	} else if e > root.E && root.Right == nil {
		root.Right = addBstNode(e)
		this.size++
		return
	}

	if e < root.E {
		this.add(this.root.Left, e)
	} else {
		this.add(this.root.Right, e)
	}
}