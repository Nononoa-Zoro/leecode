package solutions

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	reverse := reverseA2B(head, b)
	head.Next = reverseKGroup(b, k)
	return reverse
}

// reverseA2B 反转A节点到B节点
func reverseA2B(a, b *ListNode) *ListNode {
	var prev, cur *ListNode
	cur = a
	for cur != b {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
