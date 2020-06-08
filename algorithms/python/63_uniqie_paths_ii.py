from typing import List


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        row = len(obstacleGrid)
        if row == 0:
            return 0
        column = len(obstacleGrid[0])
        if column == 0:
            return 0
        dp = [0] * column
        dp[0] = 1
        for i in range(row):
            for j in range(column):
                if obstacleGrid[i][j] == 1:
                    dp[j] = 0
                elif j > 0:
                    dp[j] += dp[j - 1]
        return dp[column - 1]
