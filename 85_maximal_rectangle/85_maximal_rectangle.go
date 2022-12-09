package maximalrectangle

import "math"

func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	maxArea := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}

		maxArea = int(math.Max(float64(maxArea), float64(largestRectangleArea(heights))))
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	var stack []int
	i := 0
	res := 0

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			res = int(math.Max(float64(res), float64(width*heights[cur])))
			i -= 1
		}

		i += 1
	}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		var width int
		if len(stack) == 0 {
			width = i
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		res = int(math.Max(float64(res), float64(width*heights[cur])))
	}

	return res
}
