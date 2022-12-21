package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// iteration
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		temp := head.Next.Next
		head.Next.Next = head
		head.Next = temp
		pre = head
		head = temp
	}

	return dummy.Next
}

// recursion
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next
	head.Next = swapPairs(temp.Next)
	temp.Next = head
	return temp
}
