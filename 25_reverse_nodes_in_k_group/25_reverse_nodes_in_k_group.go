package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := ListNode{
		Val:  0,
		Next: head,
	}
	pre := &dummy
	cur := pre
	size := 0
	for cur.Next != nil {
		size += 1
		cur = cur.Next
	}

	for size >= k {
		cur = pre.Next
		for i := 1; i < k; i++ {
			temp := cur.Next
			cur.Next = temp.Next
			temp.Next = pre.Next
			pre.Next = temp
		}

		pre = cur
		size -= k
	}

	return dummy.Next
}
