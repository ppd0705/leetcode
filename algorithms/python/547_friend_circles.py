from typing import List


class Solution:
    def findCircleNum(self, M: List[List[int]]) -> int:
        count = 0
        n = len(M)
        visited = set()

        def dfs(i):
            for j in range(0, n):
                if j not in visited and M[i][j]:
                    visited.add(j)
                    dfs(j)

        for i in range(n):
            if i not in visited:
                visited.add(i)
                count += 1
                dfs(i)

        return count

    def findCircleNum2(self, M: List[List[int]]) -> int:
        count = 0
        n = len(M)
        visited = set()

        for i in range(n):
            if i not in visited:
                visited.add(i)
                queue = [i]
                count += 1
                while queue:
                    k = queue.pop(0)
                    for j in range(0, n):
                        if j not in visited and M[k][j]:
                            visited.add(j)
                            queue.append(j)

        return count

    def findCircleNum3(self, M: List[List[int]]) -> int:
        n = len(M)

        count = n
        parent = list(range(n))
        size = [1] * n

        def find(i):
            while i != parent[i]:
                parent[i] = parent[parent[i]]  # compress path
                i = parent[i]
            return i

        def union(i, j):
            nonlocal count
            pi = find(i)
            pj = find(j)
            if pi == pj:
                return

            if size[pi] > size[pj]:
                parent[pj] = pi
                size[pi] += size[pj]
            else:
                parent[pi] = pj
                size[pj] += size[pi]

            count -= 1

        for i in range(n):
            for j in range(i + 1, n):
                if M[i][j]:
                    union(i, j)

        return count


def test():
    examples = [
        ([[1, 1, 1], [1, 1, 1], [1, 1, 1]], 1),
        ([
             [1, 0, 0, 1],
             [0, 1, 1, 0],
             [0, 1, 1, 1],
             [1, 0, 1, 1]
         ], 1),
        ([[1, 1, 0],
          [1, 1, 1],
          [0, 1, 1]], 1),
        ([[1, 1, 0],
          [1, 1, 0],
          [0, 0, 1]], 2)
    ]
    s = Solution()
    for (m, expect) in examples:
        ans = s.findCircleNum3(m)
        assert ans == expect, f"expect {expect}, got {ans}"


if __name__ == "__main__":
    test()
