package smalleststringstartingfromleaf

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	path := []byte{}
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			buf := &strings.Builder{}
			n := len(path)
			for i := n - 1; i >= 0; i-- {
				buf.WriteByte(path[i])
			}

			return buf.String()
		}

		path = append(path, 'a'+byte(root.Val))
		var res string
		if root.Left != nil && root.Right != nil {
			res = min(dfs(root.Left), dfs(root.Right))
		} else if root.Left != nil {
			res = dfs(root.Left)
		} else {
			res = dfs(root.Right)
		}

		path = path[:len(path)-1]
		return res
	}

	return dfs(root)
}

func min(a, b string) string {
	if a < b {
		return a
	}

	return b
}
