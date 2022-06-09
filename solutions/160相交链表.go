package solutions

// getIntersectionNode
// 假设A链表和B链表的不相交的部分长度分别为a,b公共部分的长度为c
// A遍历完之后重新开始遍历B，总的长度为a+c+b
// B遍历完之后重新开始遍历A，总的长度为b+c+a
// 所以如果两个链表有相交节点的话，按照这种方式遍历一定会相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
