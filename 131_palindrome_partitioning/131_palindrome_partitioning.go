package palindromepartitioning

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}

	res := make([][]string, 0)
	var subRes [][]string
	for i := range s {
		if isPalindrome(s[:i+1]) {
			if i == len(s)-1 {
				res = append(res, []string{s[:i+1]})
			} else {
				subRes = partition(s[i+1:])
				for j := range subRes {
					res = append(res, append([]string{s[:i+1]}, subRes[j]...))
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
