package tree

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	rets := [][]int{}
	queue := []*TreeNode{}
	
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		ret := []int{}
		size := len(queue)
		for _, n := range queue {
			ret = append(ret, n.Val)
		}
		rets = append(rets, ret)

		for i := 0; i < size; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
	}

	return rets
}
