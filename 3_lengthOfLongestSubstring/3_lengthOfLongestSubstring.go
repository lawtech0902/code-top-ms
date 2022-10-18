/*
 * @Author: lawtech
 * @Date: 2022-05-23 14:33:21
 * @Last Modified by: lawtech
 * @Last Modified time: 2022-05-23 14:40:53
 */

package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxLen
}
