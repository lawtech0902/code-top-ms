package combinationsumiii

func combinationSum3(k, n int) [][]int {
	var (
		track []int
		res   [][]int
	)
	backTracking(n, k, 1, track, &res)
	return res
}

func backTracking(n, k, startIndex int, track []int, res *[][]int) {
	if len(track) == k {
		sum := 0
		temp := make([]int, k)
		copy(temp, track)
		for _, v := range temp {
			sum += v
		}

		if sum == n {
			*res = append(*res, temp)
		}

		return
	}

	for i := startIndex; i <= 9-(k-len(track))+1; i++ {
		track = append(track, i)
		backTracking(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
