package nextgreaterelementii

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < size*2; i++ {
		for len(stack) > 0 && nums[i%size] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = nums[i%size]
		}

		stack = append(stack, i%size)
	}

	return res
}
