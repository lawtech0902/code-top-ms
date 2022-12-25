package nextpermutation

func nextPermutation(nums []int) {
	if len(nums) < 1 {
		return
	}

	size := len(nums)
	i, j, k := size-2, size-1, size-1
	// find nums[i] < nums[j], i j相邻
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 不是最后一个排列
	if i >= 0 {
		// find nums[i] < nums[k]
		for nums[i] >= nums[k] {
			k--
		}

		// swap nums[i], nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// reverse nums[j:]
	for i, j := j, size-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
