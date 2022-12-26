package findlargestvalueineachtreerow

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	bfs(root, 0, &res)
	return res
}

func bfs(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, math.MinInt64)
	}

	(*res)[level] = max((*res)[level], root.Val)
	if root.Left != nil {
		bfs(root.Left, level+1, res)
	}

	if root.Right != nil {
		bfs(root.Right, level+1, res)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 迭代
func largestValues1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		levelMax := math.MinInt64
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			levelMax = max(node.Val, levelMax)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, levelMax)
	}

	return res
}
