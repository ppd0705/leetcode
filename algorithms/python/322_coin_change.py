from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        counts = [amount + 1] * (amount + 1)
        counts[0] = 0
        for i in range(1, amount + 1):
            for c in coins:
                if i >= c:
                    counts[i] = min(counts[i], counts[i - c] + 1)
        return -1 if counts[amount] == amount + 1 else counts[amount]

    def coinChange2(self, coins: List[int], amount: int) -> int:
        memo = {0: 0}

        def dp(n):
            if n < 0:
                return -1
            if n in memo:
                return memo[n]

            res = amount + 1

            for c in coins:
                sub = dp(n - c)
                if sub != -1 and sub < res:
                    res = sub + 1
            memo[n] = res if res < amount + 1 else -1
            return memo[n]

        return dp(amount)


def test():
    from collections import namedtuple
    sample = namedtuple("test", ["coins", "amount", "expect"])
    tests = [
        sample([1, 2, 5], 11, 3),
        sample([1], 4, 4),
        sample([2], 3, -1),
        sample([2], 5, -1),
        sample([186, 419, 83, 408], 6249, 20),
    ]
    s = Solution()
    for t in tests:
        ans = s.coinChange2(t.coins, t.amount)
        assert ans == t.expect, f"{t.amount} expect {t.expect}, got {ans}"


if __name__ == "__main__":
    test()
