package problem206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 //递归版本一
 /*
func reverseList(head *ListNode) *ListNode {
	res := head
	if head == nil {
		return head
	}
	for res.Next != nil {
		res = res.Next
	}
	recursion(head)
	head.Next = nil
	return res
}

func recursion(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	(recursion(head.Next)).Next = head
	return head
}*/

//递归二
func reverseList (head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

//迭代
func reverseList2 (head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	for p := head.Next; p != nil; {
		temp := p.Next
		p.Next = prev
		prev = p
		p = temp
	}
	head.Next = nil
	return prev
}