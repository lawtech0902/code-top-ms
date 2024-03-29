package assigncookies

import "sort"

func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	child := 0
	for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
		if s[sIdx] >= g[child] {
			child++
		}
	}

	return child
}
