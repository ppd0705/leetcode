# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head

        thead = ListNode(-1)
        thead.next = head
        current = thead
        while current.next is not None and current.next.next is not None:
            next_, next_next_ = current.next, current.next.next

            current.next = next_next_
            next_.next = next_next_.next
            current = next_next_.next = next_
        return thead.next
