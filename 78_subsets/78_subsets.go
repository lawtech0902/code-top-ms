package subsets

func subsets(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backTracking(nums, track, 0, &res)
	return res
}

func backTracking(nums, track []int, startIndex int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := startIndex; i < len(nums); i++ {
		track = append(track, nums[i])
		backTracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
