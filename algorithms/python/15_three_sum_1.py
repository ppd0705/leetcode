from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        l = len(nums)
        ret = []
        for k in range(l-2):
            if nums[k] > 0:
                break
            if k > 0 and nums[k] == nums[k-1]:
                continue
            i, j = k+1, l-1
            while i < j:
                s = nums[k] + nums[i] + nums[j]
                if s > 0:
                    j -= 1
                    while j > i and nums[j] == nums[j+1]:
                        j -= 1
                elif s < 0:
                    i += 1
                    while i < j and nums[i] == nums[i-1]:
                        i += 1
                else:
                    ret.append([nums[k], nums[i], nums[j]])
                    i += 1
                    j -= 1
                    while i < j and nums[i] == nums[i-1]:
                        i += 1
                    while j > i and nums[j] == nums[j+1]:
                        j -= 1
        return ret