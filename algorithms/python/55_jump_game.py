from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        k = len(nums)
        for i in range(len(nums)):
            if i > k:
                return False
            if k >= len(nums) - 1:
                return True

            k = max(k, i + nums[i])
        return True
