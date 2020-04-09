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

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
func maxDepth(root *TreeNode) int {
	return maxDepthNode(root, 0)
}
func maxDepthNode(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	depthLeft := maxDepthNode(root.Left, depth+1)
	depthRight := maxDepthNode(root.Right, depth+1)

	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight
}

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
// TODO:
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		valArr := []int{}
		for i := 0; i < size; i++ {
			valArr = append(valArr, queue[i].Val)
		}
		if !checkQueueIsSymmetric(valArr) {
			return false
		}
		valArr = valArr[0:0]
		for i := 0; i < size; i++ {
			node := queue[0]
			if node.Left == nil {
				queue = append(queue, &TreeNode{Val: -1})
			} else {
				queue = append(queue, node.Left)
			}
			if node.Right == nil {
				queue = append(queue, &TreeNode{Val: -1})
			} else {
				queue = append(queue, node.Left)
			}
			queue = queue[1:]
		}
	}

	return true
}

func checkQueueIsSymmetric(vals []int) bool {
	if len(vals) <= 1 {
		return true
	}
	l, r := 0, len(vals)-1
	for l < r {
		if vals[l] != vals[r] {
			return false
		}
		l++
		r--
	}
	return true
}

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	idx := -1
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break
		}
	}
	root.Left = buildTreeNode(inorder[:idx])
	root.Right = buildTreeNode(inorder[idx+1:])
	return root
}

func buildTreeNode(data []int) *TreeNode{
	if len(data) == 0 {
		return nil
	}

	if len(data) == 1 {
		return &TreeNode{Val:data[0]}
	}

	idx := (len(data)-1)/2
	if idx == 0 {
		
	}

}
