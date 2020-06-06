from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        i, j = 0, len(nums) - 1

        while i <= j:
            mid = (i + j) // 2

            if nums[mid] == target:
                return mid

            if nums[0] <= nums[mid]:
                if nums[0] <= target < nums[mid]:
                    j = mid - 1
                else:
                    i = mid + 1
            else:
                if nums[mid] < target <= nums[len(nums) - 1]:
                    i = mid + 1
                else:
                    j = mid - 1
        return -1
