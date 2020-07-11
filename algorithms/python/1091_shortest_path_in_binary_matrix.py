from typing import List
import heapq


class Solution:
    def shortestPathBinaryMatrix0(self, grid: List[List[int]]) -> int:
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

    def shortestPathBinaryMatrix2(self, grid: List[List[int]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1

        if grid[0][0] == 1 or grid[m][n] == 1:
            return -1
        from queue import PriorityQueue

        def get_distance(x: int, y: int):
            return max(m - x, n - y)

        start = (0, 0, 0, 1)
        pq = PriorityQueue()
        pq.put(start)

        grid[0][0] = 1
        while not pq.empty():
            _, x, y, step = pq.get()
            print(x, y, step)
            if x == m and y == n:
                return step
            for (dx, dy) in ((1, 1), (0, 1), (1, 0), (0, -1), (-1, 0), (-1, 1), (-1, -1), (1, -1)):
                next_x, next_y = x + dx, y + dy
                if next_x < 0 or next_x > m or next_y < 0 or next_y > n or grid[next_x][next_y] != 0:
                    continue

                pq.put((step + get_distance(next_x, next_y), next_x, next_y, step + 1))
                print(f"[{x},{y}] -->[{next_x},{next_y}]", step + 1, pq.queue)
                grid[next_x][next_y] = 1
        return -1

    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        m, n = len(grid) - 1, len(grid[0]) - 1

        if grid[0][0] == 1 or grid[m][n] == 1:
            return -1

        def heuristic(x, y):
            return max(abs(m - x), abs(n - y))

        start = (0, (0, 0, 1))
        pq = []
        heapq.heappush(pq, start)
        visited = set()  # must use a set to mark, can not replace by setting grid value = 1
        while pq:

            _, (x, y, step) = heapq.heappop(pq)
            if (x, y) in visited:
                continue
            visited.add((x, y))

            if x == m and y == n:
                return step
            for (dx, dy) in [(-1, -1), (1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (1, -1), (-1, 1)]:
                next_x, next_y = x + dx, y + dy

                if next_x < 0 or next_x > m or next_y < 0 or next_y > n or grid[next_x][next_y] != 0:
                    continue

                heapq.heappush(pq, (step + heuristic(next_x, next_y), (next_x, next_y, step + 1)))
        return -1


def test():
    s = Solution()

    for (grid, target) in [
        (
                [[0, 0, 0, 0, 1, 1, 1, 1, 0],
                 [0, 1, 1, 0, 0, 0, 0, 1, 0],
                 [0, 0, 1, 0, 0, 0, 0, 0, 0],
                 [1, 1, 0, 0, 1, 0, 0, 1, 1],
                 [0, 0, 1, 1, 1, 0, 1, 0, 1],
                 [0, 1, 0, 1, 0, 0, 0, 0, 0],
                 [0, 0, 0, 1, 0, 1, 0, 0, 0],
                 [0, 1, 0, 1, 1, 0, 0, 0, 0],
                 [0, 0, 0, 0, 0, 1, 0, 1, 0]
                 ], 11
        ),
        (
                [[0, 0, 0, 0, 1, 1],
                 [0, 1, 0, 0, 1, 0],
                 [1, 1, 0, 1, 0, 0],
                 [0, 1, 0, 0, 1, 1],
                 [0, 1, 0, 0, 0, 1],
                 [0, 0, 1, 0, 0, 0]],
                7
        ),
        (
                [[0, 0, 0], [1, 1, 0], [1, 1, 0]], 4
        )
    ]:
        ans = s.shortestPathBinaryMatrix(grid)
        assert ans == target, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
