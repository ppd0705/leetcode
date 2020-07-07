from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row = [[False] * 9 for _ in range(9)]
        column = [[False] * 9 for _ in range(9)]
        box = [[False] * 9 for _ in range(9)]

        for i in range(9):
            for j in range(9):
                if board[i][j] == ".":
                    continue
                cur = ord(board[i][j]) - 49

                k = i // 3 * 3 + j // 3
                if row[i][cur] or column[j][cur] or box[k][cur]:
                    return False
                row[i][cur] = True
                column[j][cur] = True
                box[k][cur] = True
        return True


def test():
    board = [
        ["5", "3", ".", ".", "7", ".", ".", ".", "."],
        ["6", ".", ".", "1", "9", "5", ".", ".", "."],
        [".", "9", "8", ".", ".", ".", ".", "6", "."],
        ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
        ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
        ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
        [".", "6", ".", ".", ".", ".", "2", "8", "."],
        [".", ".", ".", "4", "1", "9", ".", ".", "5"],
        [".", ".", ".", ".", "8", ".", ".", "7", "9"]
    ]
    target = True

    s = Solution()

    ans = s.isValidSudoku(board)
    assert ans == target


if __name__ == "__main__":
    test()
