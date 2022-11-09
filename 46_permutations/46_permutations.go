package permutations

func permute(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backTracking(nums, track, &res)
	return res
}

func backTracking(nums, track []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		track = append(track, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking(nums, track, res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}
