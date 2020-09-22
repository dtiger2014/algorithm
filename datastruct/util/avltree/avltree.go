package avltree

import (
	"errors"
	"fmt"
	"math"
)

type node struct {
	key    string
	value  int
	left   *node
	right  *node
	height int
}

func newNode(key string, value int) *node {
	return &node{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

type AVLTree struct {
	root *node
	size int
}

func New() *AVLTree {
	return &AVLTree{
		root: nil,
		size: 0,
	}
}

func (avl *AVLTree) GetSize() int {
	return avl.size
}

func (avl *AVLTree) IsEmpty() bool {
	return avl.size == 0
}

func (avl *AVLTree) IsBST() bool {
	keys := make([]string, 0)
	avl.inOrder(avl.root, keys)
	for i := 1; i < len(keys); i++ {
		if keys[i-1] > keys[i] {
			return false
		}
	}
	return true
}

func (avl *AVLTree) inOrder(node *node, keys []string) {
	if node == nil {
		return
	}

	avl.inOrder(node.left, keys)
	keys = append(keys, node.key)
	avl.inOrder(node.right, keys)
}

func (avl *AVLTree) IsBalanced() bool {
	return avl.isBalanced(avl.root)
}

func (avl *AVLTree) isBalanced(node *node) bool {
	if node == nil {
		return true
	}

	balanceFactor := avl.getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	return avl.isBalanced(node.left) && avl.isBalanced(node.right)
}

func (avl *AVLTree) getHeight(node *node) int {
	if node == nil {
		return 0
	}

	return node.height
}

func (avl *AVLTree) getBalanceFactor(node *node) int {
	if node == nil {
		return 0
	}
	return avl.getHeight(node.left) - avl.getHeight(node.right)
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (avl *AVLTree) leftRotate(y *node) *node {
	x := y.right
	T2 := x.left

	// 向左旋转过程
	x.left = y
	y.right = T2

	// 更新height
	y.height = int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right))) + 1)
	x.height = int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right))) + 1)
	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (avl *AVLTree) rightRotate(y *node) *node {
	x := y.right
	T2 := x.left

	// 向左旋转过程
	x.left = y
	y.right = T2

	// 更新height
	y.height = int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right))) + 1)
	x.height = int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right))) + 1)

	return x
}

func (avl *AVLTree) Add(key string, value int) {
	avl.root = avl.add(avl.root, key, value)
}

func (avl *AVLTree) add(node *node, key string, value int) *node {
	if node == nil {
		avl.size++
		return newNode(key, value)
	}

	if key < node.key {
		node.left = avl.add(node.left, key, value)
	} else if key > node.key {
		node.right = avl.add(node.right, key, value)
	} else {
		node.value = value
	}

	// 更新height
	node.height = 1 + int(math.Max(float64(avl.getHeight(node.left)), float64(avl.getHeight(node.right))))

	// 更新height
	balanceFactor := avl.getBalanceFactor(node)

	// 平衡维护
	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) >= 0 {
		return avl.rightRotate(node)
	}

	// RR
	if balanceFactor > 1 && avl.getBalanceFactor(node.right) <= 0 {
		return avl.leftRotate(node)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) < 0 {
		node.left = avl.leftRotate(node.left)
		return avl.rightRotate(node)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(node.right) > 0 {
		node.right = avl.rightRotate(node.right)
		return avl.leftRotate(node)
	}

	return node
}

func (avl *AVLTree) getNode(node *node, key string) *node {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	} else if key < node.key {
		return avl.getNode(node.left, key)
	} else {
		return avl.getNode(node.right, key)
	}
}

func (avl *AVLTree) Contains(key string) bool {
	return avl.getNode(avl.root, key) != nil
}

func (avl *AVLTree) Get(key string) (int, error) {
	node := avl.getNode(avl.root, key)
	if node == nil {
		return 0, errors.New("Not found")
	} else {
		return node.value, nil
	}
}

func (avl *AVLTree) Set(key string, newValue int) error {
	node := avl.getNode(avl.root, key)
	if node == nil {
		errStr := fmt.Sprintf("Key: %s doesn't exists", key)
		return errors.New(errStr)
	}
	node.value = newValue
	return nil
}

func (avl *AVLTree) minimum(node *node) *node {
	if node.left == nil {
		return node
	}

	return avl.minimum(node.left)
}

func (avl *AVLTree) Remove(key string) (int, error) {
	node := avl.getNode(avl.root, key)
	if node != nil {
		avl.root = avl.remove(avl.root, key)
		return node.value, nil
	}
	return 0, errors.New("remove error")
}

func (avl *AVLTree) remove(n *node, key string) *node {
	if n == nil {
		return nil
	}

	var retNode *node = nil
	if key < n.key {
		n.left = avl.remove(n.left, key)
		retNode = n
	} else if key > n.key {
		n.right = avl.remove(n.right, key)
		retNode = n
	} else {
		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			avl.size--
			retNode = rightNode
		} else if n.right == nil {
			leftNode := n.left
			n.left = nil
			avl.size--
			retNode = leftNode
		} else { // 待删除节点左右子树均不为空的情况
			// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
			// 用这个节点顶替待删除节点的位置
			successor := avl.minimum(n.right)
			successor.right = avl.remove(n.right, successor.key)
			successor.left = n.left

			n.left = nil
			n.right = nil

			retNode = successor
		}
	}
	if retNode == nil {
		return nil
	}

	// 更新height
	retNode.height = 1 + int(math.Max(float64(avl.getHeight(retNode.left)), float64(avl.getHeight(retNode.right))))

	// 计算平衡因子
	balanceFactor := avl.getBalanceFactor(retNode)

	// 平衡维护
	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) >= 0 {
		return avl.rightRotate(retNode)
	}

	// RR
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.right) <= 0 {
		return avl.leftRotate(retNode)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) < 0 {
		retNode.left = avl.leftRotate(retNode.left)
		return avl.rightRotate(retNode)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(retNode.right) > 0 {
		retNode.right = avl.rightRotate(retNode.right)
		return avl.leftRotate(retNode)
	}

	return retNode
}
