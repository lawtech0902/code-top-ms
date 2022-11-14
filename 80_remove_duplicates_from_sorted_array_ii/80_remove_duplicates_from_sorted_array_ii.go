package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	len := 0
	for _, num := range nums {
		if len < 2 || num != nums[len-2] {
			nums[len] = num
			len++
		}
	}

	return len
}
