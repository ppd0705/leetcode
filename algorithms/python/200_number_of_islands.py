from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        count = 0
        if len(grid) == 0:
            return count
        row, column = len(grid), len(grid[0])

        def dfs(x, y):
            grid[x][y] = '0'
            for r, c in [(x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)]:
                if 0 <= r < row and 0 <= c < column and grid[r][c] == '1':
                    dfs(r, c)

        for i in range(row):
            for j in range(column):
                if grid[i][j] == '1':
                    count += 1
                    dfs(i, j)
        return count

    def numIslands2(self, grid: List[List[str]]) -> int:
        count = 0
        if len(grid) == 0:
            return count
        row, column = len(grid), len(grid[0])

        from collections import deque, defaultdict
        marked = defaultdict(bool)
        queue = deque()
        for i in range(row):
            for j in range(column):
                if grid[i][j] == '1':
                    count += 1
                    queue.append((i, j))
                    marked[(i, j)] = True
                    while queue:
                        x, y = queue.popleft()
                        grid[x][y] = '0'
                        for r, c in [(x - 1, y), (x + 1, y), (x, y - 1), (x, y + 1)]:
                            if 0 <= r < row and 0 <= c < column and not marked[(r, c)] and grid[r][c] == '1':
                                queue.append((r, c))
                                marked[(r, c)] = True

        return count


def test():
    from collections import namedtuple
    sample = namedtuple("test", ["grid", "expect"])
    tests = [
        sample(
            [
                ['1', '1', '0'],
                ['0', '1', '0'],
                ['1', '1', '1'],
                ['0', '0', '0'],
            ],
            1,
        ),
        sample(
            [
                ['1', '1', '0', '0', '0'],
                ['1', '1', '0', '0', '0'],
                ['0', '0', '1', '0', '0'],
                ['0', '0', '0', '1', '1'],
            ],
            3,
        ),
        sample(
            [
                ['0', '0', '0', '0', '0'],
                ['0', '0', '0', '0', '0'],
                ['0', '0', '0', '0', '0'],
                ['0', '0', '0', '0', '0'],
            ],
            0,
        ),
    ]
    s = Solution()
    for t in tests:
        ans = s.numIslands2(t.grid)
        assert ans == t.expect, f"expect: {t.expect}, got {ans}"


if __name__ == "__main__":
    test()
