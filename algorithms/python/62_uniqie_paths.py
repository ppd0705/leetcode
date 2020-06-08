class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        def dp(i, j):
            if i < 0 or i >= m or j < 0 or j >= n:
                return 0
            if i == m - 1 or j == n - 1:
                return 1
            return dp(i + 1, j) + dp(i, j + 1)

        return dp(0, 0)

    def uniquePaths2(self, m: int, n: int) -> int:
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = 1
        for i in range(m):
            for j in range(n):
                if i > 0:
                    dp[i][j] += dp[i - 1][j]
                if j > 0:
                    dp[i][j] += dp[i][j - 1]
        return dp[m - 1][n - 1]

    def uniquePaths3(self, m: int, n: int) -> int:
        dp = [1] * n
        for i in range(1, m):
            for j in range(1, n):
                dp[j] += dp[j-1]
        return dp[n - 1]


def test():
    from collections import namedtuple
    sample = namedtuple("test", ["m", "n", "expect"])
    tests = [
        sample(3, 2, 3),
        sample(7, 3, 28),
    ]
    s = Solution()
    for t in tests:
        ans = s.uniquePaths3(t.m, t.n)
        assert ans == t.expect, f"expect {t.expect}, ans {ans}"


if __name__ == "__main__":
    test()
