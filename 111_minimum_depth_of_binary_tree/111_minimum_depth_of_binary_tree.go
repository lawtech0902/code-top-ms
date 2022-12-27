package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Right)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}

	return 1 + min(left, right)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 迭代 bfs
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left == nil && node.Right == nil {
				return level + 1
			}
		}

		level++
	}

	return level
}
