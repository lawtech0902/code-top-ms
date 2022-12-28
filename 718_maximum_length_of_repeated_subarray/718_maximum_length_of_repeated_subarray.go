package maximumlengthofrepeatedsubarray

func findLength(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
