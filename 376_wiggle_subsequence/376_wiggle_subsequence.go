package wigglesubsequence

func wiggleMaxLength(nums []int) int {
	// 上升转为下降 or 下降转为上升就计数
	res := 1
	larger := 0
	for i := 1; i < len(nums); i++ {
		if larger == 0 {
			if nums[i] != nums[i-1] {
				res++
				if nums[i] > nums[i-1] {
					larger = 1
				} else {
					larger = -1
				}
			}
		} else if larger == 1 && nums[i] < nums[i-1] {
			res++
			larger = -1
		} else if larger == -1 && nums[i] > nums[i-1] {
			res++
			larger = 1
		}
	}

	return res
}
