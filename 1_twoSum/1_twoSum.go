/*
 * @Author: lawtech
 * @Date: 2022-05-23 14:03:29
 * @Last Modified by:   lawtech
 * @Last Modified time: 2022-05-23 14:03:29
 */

package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}
