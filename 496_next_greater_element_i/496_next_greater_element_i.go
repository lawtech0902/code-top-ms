package nextgreaterelementi

func nextGreaterElement(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			if _, ok := m[nums2[top]]; ok {
				index := m[nums2[top]]
				res[index] = nums2[i]
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
