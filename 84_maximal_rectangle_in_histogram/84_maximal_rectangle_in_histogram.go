package maximalrectangleinhistogram

// 单调栈
func largestRectangleArea(heights []int) int {
	max := 0
	stack := make([]int, 0)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		for heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			temp := heights[top] * (i - left - 1)
			if temp > max {
				max = temp
			}
		}

		stack = append(stack, i)
	}

	return max
}
