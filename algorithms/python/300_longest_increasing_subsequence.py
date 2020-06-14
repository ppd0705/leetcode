from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        tails = []
        for i, n in enumerate(nums):
            if not tails or n > tails[-1]:
                tails.append(n)
            else:
                l, r = 0, len(tails)

                while l < r:
                    m = (l + r) // 2
                    if tails[m] >= n:
                        r = m
                    else:
                        l = m + 1

                tails[l] = n
        return len(tails)
