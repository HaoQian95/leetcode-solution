# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = ListNode((l1.val+l2.val)%10)
        count = (l1.val+l2.val)//10
        l1 = l1.next
        l2 = l2.next
        p = head
        while l1 and l2:
            p.next = ListNode((l1.val+l2.val+count)%10)
            count = (l1.val+l2.val+count)//10
            p = p.next
            l1 = l1.next
            l2 = l2.next
        if l2:
            while l2:
                p.next = ListNode((l2.val+count)%10)
                count = (l2.val+count)//10
                p = p.next
                l2 = l2.next
        if l1:
            while l1:
                p.next = ListNode((l1.val+count)%10)
                count = (l1.val+count)//10
                p = p.next
                l1 = l1.next
        if count!=0:
            p.next = ListNode(1)
        return head

    class Solution2(object):
        def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        re = ListNode(0)
        r = re
        carry = 0
        while l1 or l2:
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0
            r.next = ListNode((x+y+carry)%10)
            carry = (x+y+carry)//10
            r = r.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        if carry!=0:
            r.next = ListNode(1)
        return re.next