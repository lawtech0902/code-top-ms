package generateparentheses

// dfs
func generateParentheses(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, curStr string, res *[]string) {
	// 左右括号都无剩余
	if left == 0 && right == 0 {
		*res = append(*res, curStr)
	}

	// 左括号有剩余
	if left > 0 {
		dfs(left-1, right, curStr+"(", res)
	}

	// 右括号剩余大于左括号
	if right > left {
		dfs(left, right-1, curStr+")", res)
	}
}
