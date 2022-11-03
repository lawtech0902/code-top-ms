package movezeroes

func moveZeroes(nums []int) {
	count := 0
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
