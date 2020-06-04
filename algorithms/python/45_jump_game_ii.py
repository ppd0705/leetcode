from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        step, max_pos, end = 0, 0, 0

        for i in range(len(nums) - 1):
            max_pos = max(max_pos, i + nums[i])
            if i == end:
                end = max_pos
                step += 1
        return step
