# coding: utf-8
"""
链表相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

🔗 https://leetcode-cn.com/problems/add-two-numbers/
"""


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __str__(self):
        if not self.next:
            return f"{str(self.val)}"
        return str(self.val) + " -> " + self.next.__str__()


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode, carey: int = 0) -> ListNode:
        val = l1.val + l2.val + carey
        if val >= 10:
            r = ListNode(val % 10)
            carey = val // 10
        else:
            r = ListNode(val)
            carey = 0

        if l1.next and l2.next:
            r.next = self.addTwoNumbers(l1.next, l2.next, carey)
        elif l1.next and not l2.next:
            r.next = self.addTwoNumbers(l1.next, ListNode(0), carey)
        elif l2.next and not l1.next:
            r.next = self.addTwoNumbers(ListNode(0), l2.next, carey)
        else:
            if carey:
                r.next = ListNode(carey)
            else:
                r.next = None
        return r


if __name__ == '__main__':
    a = ListNode(1)
    a.next = ListNode(2)
    a.next.next = ListNode(3)

    b = ListNode(4)
    b.next = ListNode(8)
    b.next.next = ListNode(6)

    x = Solution().addTwoNumbers(a, b)
    print(x)
