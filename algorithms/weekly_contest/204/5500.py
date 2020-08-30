#  https://leetcode-cn.com/contest/weekly-contest-204/problems/maximum-length-of-subarray-with-positive-product/
from typing import List


class Solution:

    def _getMaxLen(self, nums):
        dp = [0] * len(nums)
        dp2 = [0] * len(nums)
        sign = [-1] * len(nums)

        if nums[0] == 0:
            sign[0] = 0
        elif nums[0] > 0:
            dp[0] = 1
            dp2[0] = 1
            sign[0] = 1

        for i in range(1, len(nums)):
            if nums[i] == 0:
                sign[i] = 0
            elif nums[i] < 0:
                if sign[i - 1] == -1:
                    dp[i] = dp[i - 1] + 2
                    sign[i] = 1
                else:
                    dp[i] = dp[i - 1]
            else:
                dp[i] = dp[i - 1] + 1
                if sign[i - 1] == -1:
                    sign[i] = -1
                else:
                    sign[i] = 1

                if nums[i] - 1 > 0:
                    dp2[i] = dp2[i - 1] + 1
                else:
                    dp2[i] = 1
        return max(max(dp[i] if s > 0 else 0 for i, s in enumerate(sign)), max(dp2))

    def getMaxLen(self, nums: List[int]) -> int:
        if not nums:
            return 0
        return max(self._getMaxLen(nums), self._getMaxLen(nums[::-1]))


def test():
    sl = Solution()

    for nums, target in [
        ([5, -20, -20, -39, -5, 0, 0, 0, 36, -32, 0, -7, -10, -7, 21, 20, -12, -34, 26, 2], 8),
        ([1, 3, -1, 2, 2, 6, 4], 4),
        ([-1, 2, 2], 2),
        ([-1, 2], 1),
        ([-1], 0),
        ([1, -2, -3, 4], 4),
        ([-1, -2, -3, 0, 1], 2),
        ([-1, 0], 0),
        ([1, 2, 3], 3),
        ([0, -2, 3, -1], 3),
        ([1, 2, 3, 5, -6, 4, 0, 10], 4),
    ]:
        ans = sl.getMaxLen(nums)
        assert ans == target, f"{nums}, {target}, {ans}"


if __name__ == "__main__":
    test()
