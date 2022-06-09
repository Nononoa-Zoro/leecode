package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	// 哑结点
	dummy := &ListNode{Val: -1}
	p := dummy
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	// k个链表中当前遍历最小值所在的节点
	var minNode *ListNode
	for {
		// 最小值在lists中的位置
		minIndex := -1
		for i := 0; i < n; i++ {
			if minNode == nil && lists[i] != nil {
				minNode = lists[i]
				minIndex = i
			}
			// 更新最小值node
			if lists[i] != nil && lists[i].Val < minNode.Val {
				minNode = lists[i]
				minIndex = i
			}
		}
		if minIndex != -1 {
			p.Next = minNode
			p = p.Next
			minNode = nil
			lists[minIndex] = lists[minIndex].Next
		} else {
			break
		}
	}
	return dummy.Next
}
