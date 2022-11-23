package dailytemperatures

// 单调栈（未精简）
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{0}
	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] <= temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}

			stack = append(stack, i)
		}
	}

	return res
}

// 单调栈（精简版）
func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i, v := range temperatures {
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}

		stack = append(stack, i)
	}

	return res
}
