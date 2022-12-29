package slidingwindowmaximum

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}

	// 双向队列 保存当前窗口最大值的数组位置 保证队列中数组位置的数值按从大到小排序
	queue := make([]int, 0)
	res := make([]int, 0)
	for i := range nums {
		// 保证从大到小 如果前面数小则需要依次弹出，直至满足要求
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		// 添加当前值对应的数组下标
		queue = append(queue, i)
		// 判断当前队列中队首的值是否有效
		if queue[0] <= i-k {
			queue = queue[1:]
		}

		// 当窗口长度为k时，保存当前窗口最大值
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}
