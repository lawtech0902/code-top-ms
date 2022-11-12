package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	hash1 := [26]int{}
	hash2 := [26]int{}
	for i := range s1 {
		hash1[s1[i]-'a']++
		hash2[s2[i]-'a']++
	}

	for i := len(s1); i < len(s2); i++ {
		if hash1 == hash2 {
			return true
		}

		hash2[s2[i-len(s1)]-'a']--
		hash2[s2[i]-'a']++
	}

	return hash1 == hash2
}
