package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}

			return slow
		}
	}

	return nil
}
