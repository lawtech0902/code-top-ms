package kthlargestelementinanarray

import "container/heap"

// 最小堆
func findKthLargest(nums []int, k int) int {
	ih := &IntHeap{}
	heap.Init(ih)
	for i := 0; i < k; i++ {
		heap.Push(ih, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < (*ih)[0] {
			continue
		}

		heap.Pop(ih)
		heap.Push(ih, nums[i])
	}

	return (*ih)[0]
}

type IntHeap []int // Len is the number of elements in the collection.

func (ih IntHeap) Len() int {
	return len(ih)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (ih IntHeap) Less(i int, j int) bool {
	return ih[i] < ih[j]
}

// Swap swaps the elements with indexes i and j.
func (ih IntHeap) Swap(i int, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[:n-1]
	return x
}
