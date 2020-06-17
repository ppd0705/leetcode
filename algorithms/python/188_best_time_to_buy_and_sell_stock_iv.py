from typing import List


class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        if len(prices) < 2 or k == 0:
            return 0
        if 2 * k >= len(prices):
            return self.maxProfitWithoutLimit(prices)

        dp = [[[0, 0] for _ in range(k + 1)] for _ in range(len(prices))]
        for i in range(len(prices)):
            for j in range(1, k + 1):
                if i == 0:
                    dp[i][j][1] = -prices[i]
                else:
                    dp[i][j][0] = max(dp[i - 1][j][0], dp[i - 1][j][1] + prices[i])
                    dp[i][j][1] = max(dp[i - 1][j][1], dp[i - 1][j - 1][0] - prices[i])

        return dp[len(prices) - 1][k][0]

    def maxProfitWithoutLimit(self, prices: List[int]):
        hold, not_hold = -prices[0], 0
        for p in prices[1:]:
            hold, not_hold = max(hold, not_hold - p), max(not_hold, hold + p)
        return not_hold


def test():
    tests = [
        ([2, 4, 1], 2, 2),
        ([3, 2, 6, 5, 0, 3], 2, 7),
    ]
    s = Solution()
    for (prices, k, expect) in tests:
        ans = s.maxProfit(k, prices)
        assert ans == expect, f"ans: {ans}, expect {expect}"


if __name__ == "__main__":
    test()
