from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        num_count_map = {}
        half_length = len(nums) // 2
        for n in nums:
            count = num_count_map.get(n, 0)
            if count == half_length:
                return n
            num_count_map[n] = count + 1
