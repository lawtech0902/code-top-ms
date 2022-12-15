package subarraysumequalsk

// 前缀和
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	hash := make(map[int]int)
	hash[0] = 1
	for i := range nums {
		sum += nums[i]
		if _, ok := hash[sum-k]; ok {
			count += hash[sum-k]
		}

		hash[sum] += 1
	}

	return count
}
