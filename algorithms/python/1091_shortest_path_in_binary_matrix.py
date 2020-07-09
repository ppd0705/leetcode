from typing import List


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1

        if grid[0][0] == 1 or grid[m][n] == 1:
            return -1

        queue = [(0, 0)]
        grid[0][0] = 1
        ans = 0
        offsets = ((0, 1), (0, -1), (1, 0), (-1, 0), (-1, 1), (-1, -1), (1, 1), (1, -1))
        while queue:
            l = len(queue)
            ans += 1
            for i in range(l):
                x, y = queue[i]
                if x == m and y == n:
                    return ans
                for (dx, dy) in offsets:
                    next_x, next_y = x + dx, y + dy
                    if next_x < 0 or next_x > m or next_y < 0 or next_y > n or grid[next_x][next_y] != 0:
                        continue

                    queue.append((next_x, next_y))
                    grid[next_x][next_y] = 1
            queue = queue[l:]

        return -1
