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
        if l1 is None:
            return l2
        if l2 is None:
            return l1

        n = l1.val + l2.val
        carry = 1 if n >= 10 else 0
        res = ListNode(n % 10)
        pre = res
        l1n = l1.next
        l2n = l2.next
        while (l1n is not None and l2n is not None):
            n = l1n.val + l2n.val + carry
            carry = 1 if n >= 10 else 0
            r = ListNode(n % 10)
            pre.next = r
            pre = r
            l1n = l1n.next
            l2n = l2n.next
        if l1n is None:
            if l2n is None:
                if carry == 0:
                    return res
                else:
                    last = ListNode(1)
                    pre.next = last
                    return res
            else:
                while (l2n is not None):
                    n = carry + l2n.val
                    carry = 1 if n >= 10 else 0
                    r = ListNode(n % 10)
                    pre.next = r
                    pre = r
                    l2n = l2n.next
                if carry == 0:
                    return res
                else:
                    last = ListNode(1)
                    pre.next = last
                    return res
        if l2n is None:
            while (l1n is not None):
                n = carry + l1n.val
                carry = 1 if n >= 10 else 0
                r = ListNode(n % 10)
                pre.next = r
                pre = r
                l1n = l1n.next
            if carry == 0:
                return res
            else:
                last = ListNode(1)
                pre.next = last
                return res
        
