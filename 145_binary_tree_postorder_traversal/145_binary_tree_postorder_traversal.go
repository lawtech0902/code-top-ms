package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 迭代
func postorderTraversal1(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var res []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	// reverse res
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
