package findallduplicatesinanarray

import "math"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			res = append(res, num)
		}
	}

	return res
}
