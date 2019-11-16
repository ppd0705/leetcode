# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if head is None or head.next is None:
            return head

        count = 0
        prev = None
        current = head
        while current.next is not None:
            count += 1
            next_ = current.next

            if count % 2:
                if prev is None:
                    head = next_
                else:
                    prev.next = next_

                current.next = next_.next
                next_.next = current

                prev = next_
                # current = current
            else:
                prev = current
                current = next_

        return head