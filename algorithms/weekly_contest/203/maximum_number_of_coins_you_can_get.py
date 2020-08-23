from typing import List


class Solution:
    def maxCoins(self, piles: List[int]) -> int:
        piles.sort()
        ans = 0
        cnt = len(piles) // 3
        pos = len(piles) - 2
        for i in range(cnt):
            ans += piles[pos]
            pos -= 2
        return ans


def test():
    s = Solution()

    for piles, target in [
        ([1, 2, 3], 2),
        ([2, 4, 1, 2, 7, 8], 9),
        ([9, 8, 7, 6, 5, 1, 2, 3, 4], 18),
    ]:
        ans = s.maxCoins(piles)
        assert ans == target, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
