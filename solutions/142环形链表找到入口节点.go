package solutions

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果相遇证明存在环
		// 此时slow和fast记录相遇点
		if slow == fast {
			break
		}
	}
	// 链表没有环
	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
