from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:

        def majority_element_rec(left, right):
            if left == right:
                return nums[left]

            mid = left + (right - left) // 2

            left_element = majority_element_rec(left, mid)
            right_element = majority_element_rec(mid + 1, right)

            if left_element == right_element:
                return left_element

            left_count = right_count = 0
            for n in nums[left:right + 1]:
                if n == left_element:
                    left_count += 1
                elif n == right_element:
                    right_count += 1

            return left_element if left_count > right_count else right_element

        return majority_element_rec(0, len(nums) - 1)
