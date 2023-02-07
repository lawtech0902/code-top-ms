package binarytreecameras

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	var res int
	if root == nil {
		return 0
	}

	if dfs(root, &res) == 2 {
		res++
	}

	return res
}

// 0: 该节点安装了监视器
// 1: 该节点可观，但未安装监视器
// 2: 该节点不可观
func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 1
	}

	left := dfs(node.Left, res)
	right := dfs(node.Right, res)
	if left == 2 || right == 2 {
		*res++
		return 0
	} else if left == 0 || right == 0 {
		return 1
	} else {
		return 2
	}
}
