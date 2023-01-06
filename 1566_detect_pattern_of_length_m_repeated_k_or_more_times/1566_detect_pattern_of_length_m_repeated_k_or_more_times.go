package detectpatternoflengthmrepeatedkormoretimes

func containsPattern(arr []int, m, k int) bool {
	size := len(arr)
	start := 0
	for i := m; i < size; i++ {
		if arr[i] != arr[i-m] {
			start = i - m + 1
		}

		if i-start+1 == m*k {
			return true
		}
	}

	return false
}
