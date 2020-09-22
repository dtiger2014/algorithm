package bst

import (
	"errors"
	"fmt"
	"golang-datastruct/util/arraystack"
)

type Node struct {
	Val         int
	Left, Right *Node
}

type BST struct {
	root *Node
	size int
}

func New() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

// GetSize :
func (bst *BST) GetSize() int {
	return bst.size
}

// IsEmpty :
func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

// ADD :
func (bst *BST) Add(e int) {
	bst.root = bst.add(bst.root, e)
}

// add :
func (bst *BST) add(node *Node, e int) *Node {
	if node == nil {
		bst.size++
		return &Node{
			Val:   e,
			Left:  nil,
			Right: nil,
		}
	}

	if e < node.Val {
		node.Left = bst.add(node.Left, e)
	} else if e > node.Val {
		node.Right = bst.add(node.Right, e)
	}

	return node
}

// Contains :
func (bst *BST) Contains(e int) bool {
	return bst.contains(bst.root, e)
}

func (bst *BST) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}

	if e == node.Val {
		return true
	} else if e < node.Val {
		return bst.contains(node.Left, e)
	} else if e > node.Val {
		return bst.contains(node.Right, e)
	}

	return false
}

// PreOrder : 前序遍历， 深度遍历
func (bst *BST) PreOrder() {
	bst.preOrder(bst.root)
}

func (bst *BST) preOrder(node *Node) {
	if node == nil {
		return
	}

	// 操作节点
	fmt.Println(node.Val)

	bst.preOrder(node.Left)
	bst.preOrder(node.Right)
}

// PreOrderNR : 前序遍历（非递归）
// 使用stack 完成，先push right node， 后 push left node。直到 stack 不为空
func (bst *BST) PreOrderNR() {
	if bst.root == nil {
		return
	}

	stack := arraystack.New()
	stack.Push(bst.root)

	for !stack.IsEmpty() {
		cur, _ := stack.Pop()
		fmt.Println(cur)

		// if cur.Right != nil {
		// 	stack.Push(cur.Right)
		// }
		// if cur.Left != nil {
		// 	stack.Push(cur.Left)
		// }
	}

}

// 广度优先遍历：每层进行便利， 一层一层 开始遍历。
/*
	不同于 深度便利，不会直接便利到最底层叶子节点。
	例：需要的节点，不再最底层，而是在第二层，则这时使用广度优先便利 会更好一些。减少查询量
*/

func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}

	// queue := arrayqueue.New()
	// queue.Add(bst.root)
	// for queue.IsEmpty() {
	// 	cur := queue.Remove()
	// 	fmt.Println(cur.Val)

	// 	if cur.Left != nil {
	// 		queue.Add(cur.Left)
	// 	}
	// 	if cur.Right != nil {
	// 		queue.Add(cur.Right)
	// 	}

	// }
}

// InOrder: 中序遍历
func (bst *BST) InOrder() {
	bst.inOrder(bst.root)
}

func (bst *BST) inOrder(node *Node) {
	if node == nil {
		return
	}

	bst.inOrder(node.Left)

	// 操作
	fmt.Println(node.Val)

	bst.inOrder(node.Right)
}

// PostOrder : 后序遍历
func (bst *BST) PostOrder() {
	bst.postOrder(bst.root)
}

func (bst *BST) postOrder(node *Node) {
	if node == nil {
		return
	}

	bst.postOrder(node.Left)
	bst.postOrder(node.Right)

	// 操作节点
	fmt.Println(node.Val)
}

// 最小元素
func (bst *BST) Minimum() (int, error) {
	if bst.size == 0 {
		return 0, errors.New("BST is empty")
	}

	minNode := bst.minimum(bst.root)
	return minNode.Val, nil
}

func (bst *BST) minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return bst.minimum(node.Left)
}

// 最大元素
func (bst *BST) Maximum() (int, error) {
	if bst.size == 0 {
		return 0, errors.New("BST is empty")
	}

	maxNode := bst.maximum(bst.root)
	return maxNode.Val, nil
}

func (bst *BST) maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return bst.maximum(node.Right)
}

// 从二分搜索树中删除最小值所在节点, 返回最小值
func (bst *BST) RemoveMin() int {
	ret, _ := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST) removeMin(node *Node) *Node {

	if node.Left == nil {
		rightNode := node.Right
		node.Right = nil
		bst.size--
		return rightNode
	}

	node.Left = bst.removeMin(node.Left)
	return node
}

// 从二分搜索树中删除最大值所在节点
func (bst *BST) RemoveMax() int {
	ret, _ := bst.Maximum()
	bst.root = bst.removeMax(bst.root)
	return ret
}

// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST) removeMax(node *Node) *Node {

	if node.Right == nil {
		leftNode := node.Left
		node.Left = nil
		bst.size--
		return leftNode
	}

	node.Right = bst.removeMax(node.Right)
	return node
}

// 从二分搜索树中删除元素为e的节点
func (bst *BST) Remove(e int){
	bst.root = bst.remove(bst.root, e);
}

// 删除掉以node为根的二分搜索树中值为e的节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func (bst *BST) remove(node *Node, e int) *Node{

	if node == nil {
		return nil
	}
		
	if e < node.Val {
		node.Left = bst.remove(node.Left , e)
		return node
	} else if e > node.Val {
		node.Right = bst.remove(node.Right, e);
		return node
	}else{   // e.compareTo(node.e) == 0

		// 待删除节点左子树为空的情况
		if node.Left == nil {
			rightNode := node.Right
			node.Right = nil
			bst.size--
			return rightNode
		}

		// 待删除节点右子树为空的情况
		if node.Right == nil {
			leftNode := node.Left;
			node.Left = nil
			bst.size--
			return leftNode
		}

		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := bst.minimum(node.Right)
		successor.Right = bst.removeMin(node.Right)
		successor.Left = node.Left

		node.Left = nil
		node.Right = nil

		return successor
	}
}


// Printf
func (bst *BST) String() string {
	strs := make([]byte, 0)
	strs = generateBSTString(bst.root, 0, strs)
	return string(strs)
}

func generateBSTString(node *Node, depth int, strs []byte) []byte {
	if node == nil {
		str := generateDepthString(depth) + "nil\n"
		strs = append(strs, []byte(str)...)
		return strs
	}

	str := fmt.Sprintf("%s%v\n", generateDepthString(depth), node.Val)
	strs = append(strs, []byte(str)...)

	strs = generateBSTString(node.Left, depth+1, strs)
	strs = generateBSTString(node.Right, depth+1, strs)

	return strs
}

func generateDepthString(depth int) string {
	str := ""
	for i := 0; i < depth; i++ {
		str += "--"
	}
	return str
}
