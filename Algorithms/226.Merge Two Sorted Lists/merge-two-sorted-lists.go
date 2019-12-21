package problem226

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists(head.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l1, head.Next)
	}
	return head
}