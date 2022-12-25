package predictthewinner

import "math"

// https://leetcode.cn/problems/predict-the-winner/solution/shou-hua-tu-jie-san-chong-xie-fa-di-gui-ji-yi-hua-/
// 记忆化递归
func predictTheWinner(nums []int) bool {
	size := len(nums)
	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, size)
		for j := range memo {
			memo[i][j] = math.MinInt64
		}
	}

	return helper(nums, memo, 0, size-1) >= 0
}

func helper(nums []int, memo [][]int, i, j int) int {
	if i == j {
		return nums[i]
	}

	if memo[i][j] != math.MinInt64 {
		return memo[i][j]
	}

	pickI := nums[i] - helper(nums, memo, i+1, j)
	pickJ := nums[j] - helper(nums, memo, i, j-1)
	memo[i][j] = int(math.Max(float64(pickI), float64(pickJ)))
	return memo[i][j]
}

// dp
func PredictTheWinner(nums []int) bool {
	size := len(nums)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = nums[i]
	}

	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			pickI := nums[i] - dp[i+1][j]
			pickJ := nums[j] - dp[i][j-1]
			dp[i][j] = int(math.Max(float64(pickI), float64(pickJ)))
		}
	}

	return dp[0][size-1] >= 0
}
