package nonoverlappingintervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	var res int
	// 按左边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		// 寻找交叉区间：当前区间左边界小于前一区间右边界
		if intervals[i-1][1] > intervals[i][0] {
			res++
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1]) // 覆盖区间
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
