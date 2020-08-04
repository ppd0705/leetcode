# https://leetcode-cn.com/contest/weekly-contest-199/problems/bulb-switcher-iv/

class Solution:
    def minFlips(self, target: str) -> int:
        target = int(target, base=2)
        base = 2 ** (len(bin(target)) - 2) - 1

        num = 0
        while target != 0:
            num += 1
            target = target ^ base
            base = 2 ** (len(bin(target)) - 2) - 1
        return num


def test():
    s = Solution()

    for target, num in [
        ("101", 3),
        ("0", 0),
        ("001011101", 5),
        ("10111", 3),
    ]:
        ans = s.minFlips(target)
        assert ans == num, f"expect {num}, got {ans}"


if __name__ == "__main__":
    test()
