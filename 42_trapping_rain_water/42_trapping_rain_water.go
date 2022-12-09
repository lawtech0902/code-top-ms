package trappingrainwater

// 单调栈
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	stack := make([]int, 1, len(height))
	var res int
	for i := 1; i < len(height); i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					temp := (min(height[i], height[stack[len(stack)-1]]) - height[top]) * (i - stack[len(stack)-1] - 1)
					res += temp
				}
			}

			stack = append(stack, i)
		}
	}

	return res
}

func min(a, b int) int {
	if a >= b {
		return b
	}

	return a
}

// 双指针，找左右两边最大值中小者，减去本身高度
func trap1(height []int) int {
	res, size := 0, len(height)
	if size == 0 {
		return res
	}

	leftMax, rightMax := 0, 0
	left, right := 0, size-1
	for left < right {
		if height[left] < height[right] {
			if height[left] < leftMax {
				res += leftMax - height[left]
			} else {
				leftMax = height[left]
			}

			left++
		} else {
			if height[right] < rightMax {
				res += rightMax - height[right]
			} else {
				rightMax = height[right]
			}

			right--
		}
	}

	return res
}
