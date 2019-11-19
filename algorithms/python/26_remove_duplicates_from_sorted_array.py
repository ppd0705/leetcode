from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        pos = 0
        if not nums:
            return pos

        for i in range(1, len(nums)):
            if nums[i] != nums[pos]:
                pos += 1
                nums[pos] = nums[i]
        return pos + 1
