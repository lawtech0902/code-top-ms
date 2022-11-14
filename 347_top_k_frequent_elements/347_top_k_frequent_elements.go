package topkfrequentelements

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}

	infoHeap := infoHeap{}
	heap.Init(&infoHeap)
	for num, freq := range cache {
		if len(infoHeap) < k {
			heap.Push(&infoHeap, info{
				number: num,
				freq:   freq,
			})
		} else if infoHeap[0].freq < freq {
			infoHeap[0] = info{
				number: num,
				freq:   freq,
			}
			heap.Fix(&infoHeap, 0)
		}
	}

	res := make([]int, 0, k)
	for _, info := range infoHeap {
		res = append(res, info.number)
	}

	return res
}

type info struct {
	number int
	freq   int
}

type infoHeap []info

func (ih infoHeap) Len() int {
	return len(ih)
}

// 最小堆
func (ih infoHeap) Less(i, j int) bool {
	return ih[i].freq < ih[i].freq
}

func (ih infoHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *infoHeap) Push(x interface{}) {
	*ih = append(*ih, x.(info))
}

func (ih *infoHeap) Pop() interface{} {
	x := (*ih)[len(*ih)-1]
	*ih = (*ih)[:len(*ih)-1]
	return x
}
