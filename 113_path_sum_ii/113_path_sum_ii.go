package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		res  [][]int
		path []int
	)
	dfs(root, 0, targetSum, path, &res)
	return res
}

// dfs
func dfs(root *TreeNode, level, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	if level >= len(path) {
		path = append(path, root.Val)
	} else {
		path[level] = root.Val
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		temp := make([]int, level+1)
		copy(temp, path)
		*res = append(*res, temp)
	}

	dfs(root.Left, level+1, targetSum, path, res)
	dfs(root.Right, level+1, targetSum, path, res)
}
