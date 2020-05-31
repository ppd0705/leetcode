from typing import List


class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        row = len(board)
        if row == 0:
            return board
        column = len(board[0])
        cx, cy = click[0], click[1]
        if board[cx][cy] == 'M':
            board[cx][cy] = 'X'
            return board
        elif board[cx][cy] != 'E':
            return board

        dx = [-1, -1, -1, 0, 0, 1, 1, 1]
        dy = [-1, 0, 1, -1, 1, -1, 0, 1]

        def dfs(r, c):
            count = 0
            for i in range(8):
                x, y = r + dx[i], c + dy[i]
                if 0 <= x < row and 0 <= y < column and board[x][y] == 'M':
                    count += 1

            if count > 0:
                board[r][c] = chr(ord('0') + count)
            else:
                board[r][c] = 'B'
                for i in range(8):
                    x, y = r + dx[i], c + dy[i]
                    if 0 <= x < row and 0 <= y < column and board[x][y] == 'E':
                        dfs(x, y)

        dfs(cx, cy)
        return board
