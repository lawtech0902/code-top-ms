package isMatch

func isMatch(s string, p string) bool {
	sSize, pSize := len(s), len(p)
	dp := make([][]bool, sSize+1)
	// dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	dp[0][0] = true
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || dp[i][j+1]
				}
			}
		}
	}

	return dp[sSize][pSize]
}
