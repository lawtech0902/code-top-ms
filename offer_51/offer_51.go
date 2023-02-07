package offer51

// 归并排序
func reversePairs(nums []int) int {
	count := 0
	merge(nums, 0, len(nums)-1, &count)
	return count
}

func merge(nums []int, left, right int, count *int) {
	mid := left + (right-left)>>1
	if left < right {
		merge(nums, left, mid, count)
		merge(nums, mid+1, right, count)
		mergeSort(nums, left, mid, right, count)
	}
}

func mergeSort(nums []int, left, mid, right int, count *int) {
	tempArr := make([]int, right-left+1)
	index := 0
	temp1, temp2 := left, mid+1
	for temp1 <= mid && temp2 <= right {
		if nums[temp1] <= nums[temp2] {
			tempArr[index] = nums[temp1]
			index++
			temp1++
		} else {
			// 统计逆序对个数
			*count += mid - temp1 + 1
			tempArr[index] = nums[temp2]
			index++
			temp2++
		}
	}

	// 左边剩余数移入数组
	for temp1 <= mid {
		tempArr[index] = nums[temp1]
		index++
		temp1++
	}

	// 右边剩余数移入数组
	for temp2 <= right {
		tempArr[index] = nums[temp2]
		index++
		temp2++
	}

	// 新数组覆盖nums
	for k := 0; k < len(tempArr); k++ {
		nums[k+left] = tempArr[k]
	}
}
