from typing import List


class Solution:
    def combine0(self, n: int, k: int) -> List[List[int]]:
        def dfs(start: int, option: List[int]):
            for i in range(start, n + 1 - (k - 1 - len(option))):
                option.append(i)
                if len(option) == k:
                    ans.append(option[:])
                else:
                    dfs(i + 1, option)
                option.pop()

        ans = []
        dfs(1, [])
        return ans

    def combine(self, n: int, k: int) -> List[List[int]]:
        def dfs(start: int, idx: int):
            for i in range(start, n + 1 - (k - 1 - idx)):
                option[idx] = i
                if idx == k - 1:
                    ans.append(option[:])
                else:
                    dfs(i + 1, idx + 1)

        ans = []
        option = [0] * k
        dfs(1, 0)
        return ans


def test():
    ans = Solution().combine(5, 4)
    expect = [[1, 2, 3, 4], [1, 2, 3, 5], [1, 2, 4, 5], [1, 3, 4, 5], [2, 3, 4, 5]]
    assert ans == expect, ans


if __name__ == "__main__":
    test()
