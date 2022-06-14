package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy 哑结点为了保存第n个节点的前置节点
	dummy := &ListNode{Val: -1, Next: head}
	p := head
	var length int
	for p != nil {
		length++
		p = p.Next
	}
	length = length - n
	// 从dummy开始遍历是为了找到待删除节点的前置节点pre pre->p->q
	pre := dummy
	for length > 0 {
		length--
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
