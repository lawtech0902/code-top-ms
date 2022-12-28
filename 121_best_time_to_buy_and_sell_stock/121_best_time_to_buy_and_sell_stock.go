package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < low {
			low = price
		}

		maxProfit = max(price-low, maxProfit)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
