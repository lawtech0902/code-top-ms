package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// iteration
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

// recursion
func reverseList1(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	temp := cur.Next
	cur.Next = pre
	return reverse(cur, temp)
}
