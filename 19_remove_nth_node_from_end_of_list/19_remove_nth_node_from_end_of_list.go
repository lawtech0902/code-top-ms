package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	fast := head
	slow := dummy
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}

		i++
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
