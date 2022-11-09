package restoreipaddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var (
		track []string
		res   []string
	)

	backTracking(s, 0, track, &res)
	return res
}

func backTracking(s string, startIndex int, track []string, res *[]string) {
	if startIndex == len(s) && len(track) == 4 {
		temp := strings.Join(track, ".")
		*res = append(*res, temp)
	}

	for i := startIndex; i < len(s); i++ {
		track = append(track, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(track) <= 4 && isValidIp(s, startIndex, i) {
			backTracking(s, i+1, track, res)
		} else {
			return
		}

		track = track[:len(track)-1]
	}
}

func isValidIp(s string, startIndex, end int) bool {
	val, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' {
		return false
	}

	if val > 255 {
		return false
	}

	return true
}
