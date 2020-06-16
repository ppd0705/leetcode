from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0
        dp_10 = dp_20 = 0
        dp_11 = dp_21 = -prices[0]
        for i in range(1, len(prices)):
            dp_20 = max(dp_20, dp_21 + prices[i])
            dp_21 = max(dp_21, dp_10 - prices[i])
            dp_10 = max(dp_10, dp_11 + prices[i])
            dp_11 = max(dp_11, -prices[i])
        return dp_20

    def maxProfit2(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0
        n = 2
        dp = [[[0, 0] for _ in range(n + 1)] for _ in range(len(prices))]

        for i in range(0, len(prices)):
            for j in range(1, n + 1, 1):
                if i == 0:
                    dp[i][j][0] = 0
                    dp[i][j][1] = -prices[i]
                    continue
                dp[i][j][0] = max(dp[i - 1][j][0], dp[i - 1][j][1] + prices[i])
                dp[i][j][1] = max(dp[i - 1][j][1], dp[i - 1][j - 1][0] - prices[i])
            # print(i, prices[i], [str(j) + ": " + str(dp[i][j]) for j in range(n + 1)])
        return dp[len(prices) - 1][n][0]


def test():
    tests = [
        ([3, 3, 5, 0, 0, 3, 1, 4], 6),
    ]
    s = Solution()
    for (prices, expect) in tests:
        ans = s.maxProfit(prices)
        assert ans == expect, f"expect {expect}, got {ans}"


if __name__ == "__main__":
    test()
