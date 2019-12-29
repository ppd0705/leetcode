from typing import List


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        def dfs(queens, xy_sum, xy_diff):
            row = len(queens)
            if row == n:
                solutions.append(queens)
                return

            for column in range(n):
                if column not in queens and (row + column not in xy_sum) and (row - column not in xy_diff):
                    dfs(queens + [column], xy_sum + [row + column], xy_diff + [row - column])

        solutions = []
        dfs([], [], [])
        return [['.' * c + 'Q' + '.' * (n - c - 1) for c in s] for s in solutions]
