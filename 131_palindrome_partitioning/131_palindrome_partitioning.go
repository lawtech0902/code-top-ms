package palindromepartitioning

func partition(s string) [][]string {
	var (
		tempString []string
		res        [][]string
	)

	backTracking(s, tempString, 0, &res)
	return res
}

func backTracking(s string, tempString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) {
		t := make([]string, len(tempString))
		copy(t, tempString)
		*res = append(*res, t)
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			tempString = append(tempString, s[startIndex:i+1])
		} else {
			continue
		}

		backTracking(s, tempString, i+1, res)
		tempString = tempString[:len(tempString)-1]
	}
}

func isPalindrome(s string, startIndex, end int) bool {
	for startIndex < end {
		if s[startIndex] != s[end] {
			return false
		}

		startIndex++
		end--
	}

	return true
}
