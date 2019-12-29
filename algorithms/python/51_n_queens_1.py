from typing import List


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        def is_valid(row, column):
            return not cols[column] and not hill_diagonals[row + column] and not dale_diagonals[row - column]

        def place_queen(row, column):
            cols[column] = 1
            hill_diagonals[row + column] = 1
            dale_diagonals[row - column] = 1
            queens[row] = column

        def add_solution():
            pass

        def remove_queen(row, column):
            cols[column] = 0
            hill_diagonals[row + column] = 0
            dale_diagonals[row - column] = 0
            queens[row] = -1

        def add_solution():
            solution = []
            for row, column in enumerate(queens):
                solution.append('.' * column + 'Q' + '.' * (n - column - 1))
            output.append(solution)

        def backtrack(row=0):
            for column in range(n):
                if is_valid(row, column):
                    place_queen(row, column)
                    if row + 1 == n:
                        add_solution()
                    else:
                        backtrack(row + 1)
                    remove_queen(row, column)

        cols = [0] * n
        hill_diagonals = [0] * (2 * n - 1)
        dale_diagonals = [0] * (2 * n - 1)
        queens = [-1] * n
        output = []
        backtrack()
        return output


print(Solution().solveNQueens(8))