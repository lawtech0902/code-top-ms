package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 迭代
func inorderTraversal1(root *TreeNode) []int {
	var (
		stack []*TreeNode
		res   []int
	)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		size := len(stack)
		if size == 0 {
			return res
		}

		node := stack[size-1]
		stack = stack[:size-1]
		res = append(res, node.Val)
		root = node.Right
	}
}
