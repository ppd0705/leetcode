from typing import List
from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        if k == 1:
            return nums

        ret = []
        q = deque()

        def deque_push(v):
            while q and v > q[-1]:
                q.pop()
            q.append(v)

        def deque_pop(v):
            if q and q[0] == v:
                q.popleft()

        for i, n in enumerate(nums):
            if i < k - 1:
                deque_push(n)
            else:
                deque_push(n)
                ret.append(q[0])
                deque_pop(nums[i - k + 1])

        return ret
