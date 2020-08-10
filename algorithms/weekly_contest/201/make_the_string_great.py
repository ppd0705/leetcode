# https://leetcode-cn.com/contest/weekly-contest-201/problems/make-the-string-great/


class Solution:
    def makeGood(self, s: str) -> str:
        i = 1
        while i <= len(s) - 1:
            if ord(s[i]) - ord(s[i - 1]) in (-32, 32):
                s = s[:i - 1] + s[i + 1:]
                i -= 2
            i += 1
            if i == 0:
                i = 1
        return s


def test():
    sl = Solution()
    for s, target in [
        ("abBAcC", ""),
        ("mC", "mC"),
        ("leEeetcode", "leetcode"),
        ("s", "s"),
        ("ss", "ss"),
        ("aBbcCA", ""),
    ]:
        ans = sl.makeGood(s)
        assert ans == target, f"target {target},  ans: {ans}"


if __name__ == "__main__":
    test()
