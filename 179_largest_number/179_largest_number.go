package largestnumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	size := len(nums)
	arr := make([]string, size)
	for i, num := range nums {
		arr[i] = strconv.Itoa(num)
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i]+arr[j] > arr[j]+arr[i]
		}

		return arr[i] > arr[j]
	})

	res := strings.Join(arr, "")
	if res[0] == '0' {
		return "0"
	}

	return res
}
