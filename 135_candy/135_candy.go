package candy

func candy(ratings []int) int {
	size := len(ratings)
	candies := make([]int, size)
	sum := 0
	for i := range candies {
		candies[i] = 1
	}

	// 从左向右比较
	for i := 0; i < size-1; i++ {
		if ratings[i] < ratings[i+1] {
			candies[i+1] = candies[i] + 1
		}
	}

	// 从右向左比较，去两者最大值
	for i := size - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candies[i-1] = max(candies[i]+1, candies[i-1])
		}
	}

	for _, candy := range candies {
		sum += candy
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
