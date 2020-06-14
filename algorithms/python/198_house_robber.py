from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        if not nums:
            return 0
        dp = nums.copy()

        for i in range(2, len(nums)):
            dp[i] += max(dp[:i - 1])
        return max(dp[-2:])

    def rob2(self, nums: List[int]) -> int:
        if not nums:
            return 0
        if len(nums) == 1:
            return nums[0]
        dp = [0] * len(nums)
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])
        for i in range(2, len(nums)):
            dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
        return dp[-1]

    def rob3(self, nums: List[int]) -> int:
        if not nums:
            return 0
        if len(nums) == 1:
            return nums[0]
        first = nums[0]
        second = max(nums[0], nums[1])
        for i in range(2, len(nums)):
            first, second = second, max(second, first + nums[i])
        return second


def test():
    tests = [
        ([100, 0, 1, 300, 3], 400),
        ([100], 100),
        ([], 0),
    ]
    s = Solution()
    for (nums, expect) in tests:
        ans = s.rob3(nums)
        assert ans == expect, f"expect {expect}, got {ans}"


if __name__ == "__main__":
    test()
