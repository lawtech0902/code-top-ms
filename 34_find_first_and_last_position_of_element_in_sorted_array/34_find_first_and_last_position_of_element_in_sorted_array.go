package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	leftBorder := getLeftBorder(nums, target)
	rightBorder := getRightBorder(nums, target)
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}

	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}

	return []int{-1, -1}
}

func getLeftBorder(nums []int, target int) int {
	l, r := 0, len(nums)-1
	border := -2
	for l <= r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m - 1
			border = r
		} else {
			l = m + 1
		}
	}

	return border
}

func getRightBorder(nums []int, target int) int {
	l, r := 0, len(nums)-1
	border := -2
	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
			border = l
		}
	}

	return border
}
