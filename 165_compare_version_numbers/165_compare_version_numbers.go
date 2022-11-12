package compareversionnumbers

func compareVersion(version1, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		a, b := 0, 0
		for i < len(version1) && version1[i] != '.' {
			a = int(version1[i]-'0') + a*10
			i++
		}

		for j < len(version2) && version2[j] != '.' {
			b = int(version2[j]-'0') + b*10
			j++
		}

		if a > b {
			return 1
		} else if a < b {
			return -1
		}

		i++
		j++
	}

	return 0
}
