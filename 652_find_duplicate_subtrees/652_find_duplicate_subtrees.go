package findduplicatesubtrees

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int)
	res := make([]*TreeNode, 0)
	collect(root, count, &res)
	return res
}

func collect(node *TreeNode, count map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return "#"
	}

	serial := strconv.Itoa(node.Val) + "," + collect(node.Left, count, res) + "," + collect(node.Right, count, res)
	count[serial]++
	if count[serial] == 2 {
		*res = append(*res, node)
	}

	return serial
}
