# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head
        else:
            tail = self.reverseList(head.next)
            prev = head.next
            prev.next = head
            head.next = None
            return tail