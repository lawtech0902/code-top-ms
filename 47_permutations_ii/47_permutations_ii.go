package permutationsii

import "sort"

// 使用history
func permuteUnique(nums []int) [][]int {
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

	history := [21]int{}
	for i := 0; i < len(nums); i++ {
		if history[nums[i]+10] == 1 {
			continue
		}

		cur := nums[i]
		track = append(track, cur)
		history[nums[i]+10] = 1
		nums = append(nums[:i], nums[i+1:]...)
		backTracking(nums, track, res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}

// 排序后判断
func permuteUnique1(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	sort.Ints(nums)
	backTracking1(nums, track, &res)
	return res
}

func backTracking1(nums, track []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		cur := nums[i]
		track = append(track, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking1(nums, track, res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}
