package mergeksortedlists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func merge(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 1 {
		return lists[0]
	}

	if size == 2 {
		var (
			l1, l2   = lists[0], lists[1]
			res, cur *ListNode
		)

		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}

		if l1.Val < l2.Val {
			res, cur, l1 = l1, l1, l1.Next
		} else {
			res, cur, l2 = l2, l2, l2.Next
		}

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next, l1 = l1, l1.Next
			} else {
				cur.Next, l2 = l2, l2.Next
			}

			cur = cur.Next
		}

		if l1 != nil {
			cur.Next = l1
		}

		if l2 != nil {
			cur.Next = l2
		}

		return res
	}

	half := size / 2
	return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
}

// 构造最小堆
// Len is the number of elements in the collection.
func (h minHeap) Len() int {
	return len(h)
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
func (h minHeap) Less(i int, j int) bool {
	if h[i] == nil {
		return false
	}

	if h[j] == nil {
		return true
	}

	return h[i].Val < h[j].Val
}

// Swap swaps the elements with indexes i and j.
func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() any {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

type minHeap []*ListNode

func mergeKLists1(lists []*ListNode) *ListNode {
	temp := minHeap(lists)
	h := &temp
	heap.Init(h)
	var head, last *ListNode
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		if v == nil {
			continue
		}

		if last != nil {
			last.Next = v
		}

		if head == nil {
			head = v
		}

		last = v
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}

	return head
}
