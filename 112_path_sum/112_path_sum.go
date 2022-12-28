package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 迭代
type Pair struct {
	Node    *TreeNode
	PathSum int
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []Pair{{root, root.Val}}
	for len(stack) != 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.Node.Left == nil && pair.Node.Right == nil && pair.PathSum == targetSum {
			return true
		}

		if pair.Node.Right != nil {
			stack = append(stack, Pair{pair.Node.Right, pair.PathSum + pair.Node.Right.Val})
		}

		if pair.Node.Left != nil {
			stack = append(stack, Pair{pair.Node.Left, pair.PathSum + pair.Node.Left.Val})
		}
	}

	return false
}
