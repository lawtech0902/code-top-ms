/*
 * @Author: lawtech
 * @Date: 2022-05-23 14:48:23
 * @Last Modified by: lawtech
 * @Last Modified time: 2022-05-24 14:25:33
 */

package findmediansortedarrays

import "math"

// O(log (m+n)), 二分
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	return (float64(findKth(nums1, nums2, 0, 0, left)) + float64(findKth(nums1, nums2, 0, 0, right))) / 2.00
}

// i: nums1的起始位置 j: nums2的起始位置
func findKth(nums1, nums2 []int, i, j, k int) int {
	// nums1 空数组
	if i >= len(nums1) {
		return nums2[j+k-1]
	}

	// nums2 空数组
	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	midVal1, midVal2 := math.MaxInt32, math.MaxInt32
	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	}

	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	}

	if midVal1 < midVal2 {
		return findKth(nums1, nums2, i+k/2, j, k-k/2)
	} else {
		return findKth(nums1, nums2, i, j+k/2, k-k/2)
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}
