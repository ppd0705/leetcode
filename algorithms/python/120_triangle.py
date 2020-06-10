from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        dp = triangle[-1].copy()
        for i in range(len(triangle) - 2, -1, -1):
            for j in range(len(triangle[i])):
                dp[j] = triangle[i][j] + min(dp[j], dp[j + 1])
        return dp[0]


def test():
    triangle = [[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]]
    s = Solution()
    ans = s.minimumTotal(triangle)
    assert ans == 11, f"got {ans}"


if __name__ == "__main__":
    test()
