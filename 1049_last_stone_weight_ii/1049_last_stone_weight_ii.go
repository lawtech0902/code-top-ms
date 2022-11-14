package laststoneweightii

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}

	target := sum / 2
	// 1 <= stones.length <= 30，1 <= stones[i] <= 1000
	// 30 * 1000 / 2 + 1 = 15001
	dp := make([]int, 15001)
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}

	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
