package combinations

func combine(n int, k int) [][]int {
	var (
		res    [][]int
		subRes []int
	)
	if n <= 0 || k <= 0 || k > n {
		return res
	}

	backTracking(n, k, 1, subRes, &res)
	return res
}

func backTracking(n, k, start int, subRes []int, res *[][]int) {
	if len(subRes) == k {
		temp := make([]int, k)
		copy(temp, subRes)
		*res = append(*res, temp)
	}

	if len(subRes)+n-start+1 < k {
		return
	}

	for i := start; i <= n; i++ {
		subRes = append(subRes, i)
		backTracking(n, k, i+1, subRes, res)
		subRes = subRes[:len(subRes)-1]
	}
}
