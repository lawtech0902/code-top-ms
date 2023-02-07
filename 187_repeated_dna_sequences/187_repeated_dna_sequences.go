package repeateddnasequences

func findRepeatedDnaSequences(s string) []string {
	var res []string
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		subStr := s[i : i+10]
		if _, exists := m[subStr]; exists {
			m[subStr] += 1
		} else {
			m[subStr] = 1
		}
	}

	for subStr, cnt := range m {
		if cnt > 1 {
			res = append(res, subStr)
		}
	}

	return res
}

// 结合位运算
func findRepeatedDnaSequences1(s string) []string {
	var (
		res            []string
		hash, accessed [1 << 20]bool
		codex          = map[byte]int{
			'A': 0,
			'C': 1,
			'G': 2,
			'T': 3,
		}
	)

	const MASK = 0x0fffff
	var code int
	for i := range s {
		code = (code<<2 | codex[s[i]]) & MASK
		if i < 9 {
			continue
		}

		if hash[code] {
			if !accessed[code] {
				res = append(res, s[i-9:i+1])
				accessed[code] = true
			}
		}

		hash[code] = true
	}

	return res
}
