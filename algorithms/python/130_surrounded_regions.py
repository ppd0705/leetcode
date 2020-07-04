from typing import List


class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """

        m = len(board)
        if m < 3:
            return

        n = len(board[0])
        if n < 3:
            return

        def dfs(i, j):
            for (dx, dy) in ((0, 1), (0, -1), (1, 0), (-1, 0)):
                x, y = i + dx, j + dy

                if 0 <= x < m and 0 <= y < n and board[x][y] == "O":
                    board[x][y] = "B"
                    dfs(x, y)

        for i in range(m):
            for j in range(n):
                if (i == 0 or j == 0 or i == m - 1 or j == n - 1) and board[i][j] == "O":
                    board[i][j] = "B"
                    dfs(i, j)
        for i in range(m):
            for j in range(n):
                if board[i][j] == "O":
                    board[i][j] = "X"
                elif board[i][j] == "B":
                    board[i][j] = "O"


def test():
    s = Solution()
    for board, expect in [
        ([["X", "X", "X", "X"], ["X", "O", "O", "X"], ["X", "X", "O", "X"], ["X", "O", "X", "X"]],
         [["X", "X", "X", "X"], ["X", "X", "X", "X"], ["X", "X", "X", "X"], ["X", "O", "X", "X"]]),
        ([["O", "O", "O"], ["O", "O", "O"], ["O", "O", "O"]], [["O", "O", "O"], ["O", "O", "O"], ["O", "O", "O"]]),
    ]:
        s.solve(board)
        assert board == expect, f"board {board}\nexpect {expect}\n"


if __name__ == "__main__":
    test()
