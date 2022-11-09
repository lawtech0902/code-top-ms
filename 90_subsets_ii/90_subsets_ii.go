package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	sort.Ints(nums)
	backTracking(nums, track, 0, &res)
	return res
}

func backTracking(nums, track []int, startIndex int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backTracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
