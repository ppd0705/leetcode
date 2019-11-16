from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        ret = set()
        l = len(nums)
        for k in range(l - 2):
            x = nums[k]
            cache = {}
            for i in range(k + 1, l):
                y = nums[i]
                if - x - y in cache:
                    ret.add([x, -x - y, y])
                else:
                    cache[y] = 1
        return list(ret)
