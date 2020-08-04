# https://leetcode-cn.com/contest/weekly-contest-198/problems/water-bottles/
class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        ans = numBottles
        numEmpty = numBottles
        while numEmpty >= numExchange:
            div, mod = divmod(numEmpty, numExchange)
            ans += div
            numEmpty = div + mod
        return ans


def test():
    s = Solution()
    for (numBottles, numExchange, target) in [
        (9, 3, 13),
        (4, 3, 5),
        (15, 4, 19),
        (5, 5, 6),
        (2, 3, 2),
    ]:
        ans = s.numWaterBottles(numBottles, numExchange)
        assert ans == target, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
