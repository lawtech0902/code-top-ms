package binarytreepreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	var (
		res []int
		dfs func(root *TreeNode)
	)

	dfs = func(root *TreeNode) {
		if root != nil {
			res = append(res, root.Val)
			dfs(root.Left)
			dfs(root.Right)
		}
	}

	dfs(root)
	return res
}

// 迭代
func preorderTraversal1(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var res []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return res
}
