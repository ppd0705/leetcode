from typing import List


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = grid[0].copy()
        for i in range(1, n):
            dp[i] += dp[i - 1]

        for i in range(1, m):
            for j in range(0, n):
                if j == 0:
                    dp[j] += grid[i][j]
                else:
                    dp[j] = min(dp[j], dp[j - 1]) + grid[i][j]

        return dp[-1]


def test():
    tests = [
        (
            [
                [1, 3, 1],
                [1, 5, 1],
                [4, 2, 1]
            ], 7
        )
    ]
    s = Solution()
    for grid, target in tests:
        ans = s.minPathSum(grid)
        assert ans == target, f"except {target}, got {ans}"


if __name__ == "__main__":
    test()
