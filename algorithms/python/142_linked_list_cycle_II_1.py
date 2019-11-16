# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def detectCycle(self, head: ListNode) -> ListNode:
        cache = set()
        while head is not None:
            if head in cache:
                return head
            else:
                cache.add(head)
                head = head.next
        return None
