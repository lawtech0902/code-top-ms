package binarytreelevelordertraversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func levelOrder(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(*TreeNode, int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}

// 迭代
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	var tempArr []int
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			tempArr = append(tempArr, node.Val)
		}

		res = append(res, tempArr)
		tempArr = []int{}
	}

	return res
}
