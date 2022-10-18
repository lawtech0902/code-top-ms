package myAtoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " \t")
	res, flag, size, idx := 0, 1, len(s), 0
	if idx < size {
		if s[idx] == '+' {
			flag = 1
			idx++
		} else if s[idx] == '-' {
			flag = -1
			idx++
		}
	}

	for idx < size && s[idx] >= '0' && s[idx] <= '9' {
		res = res*10 + int(s[idx]) - '0'
		if flag*res > math.MaxInt32 {
			return math.MaxInt32
		} else if flag*res < math.MaxInt32 {
			return math.MinInt32
		}

		idx++
	}

	return flag * res
}
