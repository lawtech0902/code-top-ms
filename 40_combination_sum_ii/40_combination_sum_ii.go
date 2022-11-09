package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)

	sort.Ints(candidates)
	backTracking(target, 0, 0, candidates, track, &res)
	return res
}

func backTracking(target, sum, startIndex int, candidates, track []int, res *[][]int) {
	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// i > startIndex: 同一树层
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		sum += candidates[i]
		backTracking(target, sum, i+1, candidates, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}
