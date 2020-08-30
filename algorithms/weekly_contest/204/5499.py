# https://leetcode-cn.com/contest/weekly-contest-204/problems/detect-pattern-of-length-m-repeated-k-or-more-times/

from typing import List


class Solution:
    def containsPattern(self, arr: List[int], m: int, k: int) -> bool:
        i = 0
        while i <= len(arr) - m * k:
            valid = True
            for j in range(i, i + m):
                v = arr[j]
                for l in range(1, k):
                    if arr[j + l * m] != v:
                        valid = False
                        break
                if not valid:
                    i = j
                    break
            else:
                return True
            i += 1
        return False


def test():
    sl = Solution()
    for arr, m, k, target in [
        ([3, 6, 6, 6, 5, 1, 5, 2, 2, 3, 1, 5, 2, 6, 1, 5, 1, 2, 6, 3, 3, 5, 3, 6, 3, 4], 6, 2, False),
        ([1, 1, 1], 1, 3, True),
        ([1, 2, 3, 1, 2], 2, 2, False),
        ([2, 2, 2, 2], 2, 3, False),
        ([1, 2, 1, 2, 1, 1, 1, 3], 2, 2, True),
    ]:
        ans = sl.containsPattern(arr, m, k)
        assert ans is target, f"{arr}, {m}, {k}, {ans}, {target}"


if __name__ == "__main__":
    test()
