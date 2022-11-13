package maximumsubarray

import "math"

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		res = int(math.Max(float64(res), float64(sum)))
	}

	return res
}
