package maximumdepthofbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}

// 迭代 bfs
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return level
}
