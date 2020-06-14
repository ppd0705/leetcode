from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        def _rob(nums: List[int]) -> int:
            if not nums:
                return 0
            if len(nums) == 1:
                return nums[0]
            first = nums[0]
            second = max(nums[0], nums[1])
            for i in range(2, len(nums)):
                first, second = second, max(second, first + nums[i])
            return second
        if len(nums) == 1:
            return nums[0]
        return max(_rob(nums[1:]), _rob(nums[:-1]))


def test():
    tests = [
        ([100, 0, 1, 300], 300),
        ([100, 0, 1, 300, 3], 400),
        ([100], 100),
        ([], 0),
    ]
    s = Solution()
    for (nums, expect) in tests:
        ans = s.rob(nums)
        assert ans == expect, f"expect {expect}, got {ans}"


if __name__ == "__main__":
    test()
