package maximumwidthramp

func maxWidthRamp(nums []int) int {
	res := 0
	size := len(nums)
	stack := []int{0}
	for i := 0; i < size; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i) // 严格单调递减栈
		}
	}

	for j := size - 1; j >= res; j-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[j] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = max(res, j-index)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
