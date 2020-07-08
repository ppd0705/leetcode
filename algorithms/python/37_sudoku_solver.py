from typing import List


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """

        row = [[False for _ in range(9)] for _ in range(9)]
        column = [[False for _ in range(9)] for _ in range(9)]
        box = [[False for _ in range(9)] for _ in range(9)]

        to_fill = [[board[i][j] == "." for j in range(9)] for i in range(9)]

        missed = 0
        for i in range(9):
            for j in range(9):
                if board[i][j] != ".":
                    k = i // 3 * 3 + j // 3
                    n = ord(board[i][j]) - 49
                    row[i][n] = True
                    column[j][n] = True
                    box[k][n] = True
                else:
                    missed += 1

        def can_fill(i, j, v):
            return not row[i][v] and not column[j][v] and not box[i // 3 * 3 + j // 3][v]

        def dfs(i, j, cnt):
            if cnt == 0:
                # all done
                return 0

            if to_fill[i][j]:
                for l in range(0, 9):
                    if can_fill(i, j, l):
                        board[i][j] = chr(l + 49)
                        row[i][l] = True
                        column[j][l] = True
                        box[i // 3 * 3 + j // 3][l] = True

                        if (dfs(i + 1, 0, cnt - 1) if j == 8 else dfs(i, j + 1, cnt - 1)) == 0:
                            # all done
                            return 0

                        row[i][l] = False
                        column[j][l] = False
                        box[i // 3 * 3 + j // 3][l] = False
                        board[i][j] = "."
                # no resolution
                return -1
            else:
                if j == 8:
                    return dfs(i + 1, 0, cnt)
                else:
                    return dfs(i, j + 1, cnt)

        dfs(0, 0, missed)


if __name__ == "__main__":
    board = [["5", "3", ".", ".", "7", ".", ".", ".", "."],
             ["6", ".", ".", "1", "9", "5", ".", ".", "."],
             [".", "9", "8", ".", ".", ".", ".", "6", "."],
             ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
             ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
             ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
             [".", "6", ".", ".", ".", ".", "2", "8", "."],
             [".", ".", ".", "4", "1", "9", ".", ".", "5"],
             [".", ".", ".", ".", "8", ".", ".", "7", "9"]]

    target = [['5', '3', '4', '6', '7', '8', '9', '1', '2'],
         ['6', '7', '2', '1', '9', '5', '3', '4', '8'],
         ['1', '9', '8', '3', '4', '2', '5', '6', '7'],
         ['8', '5', '9', '7', '6', '1', '4', '2', '3'],
         ['4', '2', '6', '8', '5', '3', '7', '9', '1'],
         ['7', '1', '3', '9', '2', '4', '8', '5', '6'],
         ['9', '6', '1', '5', '3', '7', '2', '8', '4'],
         ['2', '8', '7', '4', '1', '9', '6', '3', '5'],
         ['3', '4', '5', '2', '8', '6', '1', '7', '9']]

    s = Solution()
    s.solveSudoku(board)

    assert board == target
