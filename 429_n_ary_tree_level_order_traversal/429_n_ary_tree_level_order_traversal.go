package narytreelevelordertraversal

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	bfs(root, 0, &res)
	return res
}

func bfs(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		bfs(v, level+1, res)
	}

	for len(*res) <= level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
}

// 迭代
func levelOrder1(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*Node{root}
	for len(queue) != 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}

		res = append(res, level)
	}

	return res
}
