package reconstructitinerary

import "sort"

type pair struct {
	target  string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	// 字典序
	return p[i].target < p[j].target
}

func findItinerary(tickets [][]string) []string {
	var res []string
	res = append(res, "JFK")

	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}

		targets[ticket[0]] = append(targets[ticket[0]], &pair{
			target:  ticket[1],
			visited: false,
		})
	}

	for _, v := range targets {
		sort.Sort(v)
	}

	backTracking(tickets, targets, &res)
	return res
}

func backTracking(tickets [][]string, targets map[string]pairs, res *[]string) bool {
	if len(*res) == len(tickets)+1 {
		return true
	}

	for _, pair := range targets[(*res)[len(*res)-1]] {
		if pair.visited == false {
			*res = append(*res, pair.target)
			pair.visited = true
			if backTracking(tickets, targets, res) {
				return true
			}

			pair.visited = false
			*res = (*res)[:len(*res)-1]
		}
	}

	return false
}
