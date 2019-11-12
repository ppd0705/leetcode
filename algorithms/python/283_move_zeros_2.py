class Solution(object):
    def moveZeroes(self, nums: list) -> None:
        # i为0的索引, j为非0的索引
        i = j = 0
        while j < len(nums):
            if nums[j] != 0:
                if i != j:
                    nums[i] = nums[j]
                    nums[j] = 0
                i += 1
            j += 1