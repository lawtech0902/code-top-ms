package maxarea

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res, area := 0, 0
	for l < r {
		if height[l] < height[r] {
			area = height[l] * (r - l)
			l += 1
		} else {
			area = height[r] * (r - l)
			r -= 1
		}

		res = int(math.Max(float64(area), float64(res)))
	}

	return res
}
