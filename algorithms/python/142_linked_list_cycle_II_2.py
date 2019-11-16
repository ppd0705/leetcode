# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


# a + b = l
# f = 2s
# f = s + nb
# s = nb

class Solution(object):
    def detectCycle(self, head: ListNode) -> ListNode:
        fast = slow = head

        while True:
            if fast is None or fast.next is None:
                return None
            fast = fast.next.next
            slow = slow.next
            if fast is slow:
                break

        fast = head
        while fast is not slow:
            fast = fast.next
            slow = slow.next

        return fast
