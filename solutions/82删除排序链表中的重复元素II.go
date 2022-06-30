package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 哑结点
	dummy := &ListNode{Val: -1, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			p := cur.Next
			for p != nil && p.Next != nil && p.Val == p.Next.Val {
				p = p.Next
			}
			if p != nil {
				cur.Next = p.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
