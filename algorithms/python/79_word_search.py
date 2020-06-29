from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        row = len(board)
        column = len(board[0])

        visted = [[0 for _ in range(column)] for _ in range(row)]

        offsets = [(0, 1), (0, -1), (1, 0), (-1, 0)]

        def dfs(i: int, j: int, level: int) -> bool:
            if visted[i][j] or board[i][j] != word[level]:
                return False
            if level == len(word) - 1:
                return True

            visted[i][j] = 1
            for (x_, y_) in offsets:
                x, y = i + x_, j + y_
                if 0 <= x < row and 0 <= y < column and dfs(x, y, level + 1):
                    return True

            visted[i][j] = 0
            return False

        for r in range(row):
            for c in range(column):
                if dfs(r, c, 0):
                    return True

        return False


def test():
    board = [
        ['A', 'B', 'C', 'E'],
        ['S', 'F', 'C', 'S'],
        ['A', 'D', 'E', 'E']
    ]
    s = Solution()
    assert s.exist(board, "ABCCED")
    assert s.exist(board, "SEE")
    assert not s.exist(board, "ABCB")


if __name__ == "__main__":
    test()
