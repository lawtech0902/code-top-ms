package binarytreemaximumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if nil == root {
		return 0
	}
	l := dfs(root.Left, maxSum)  // 左子树到达root的最大路径和l
	r := dfs(root.Right, maxSum) //右子树到达root的最大路径和r
	// 最大和为 l+r+root,l+root,r+root,root的最大值
	*maxSum = max(*maxSum, max(l+r+root.Val, max(max(l, r)+root.Val, root.Val)))
	return max(max(l, r)+root.Val, root.Val) // 返回的路径最大和为 l+root,r+root,root的最大值，不能带l+r+root，因为l+r+root是闭环路径
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
