class Solution:
    def numDecodings0(self, s: str) -> int:
        if len(s) == 0 or s[0] == "0":
            return 0
        dp = [0] * (len(s) + 1)
        dp[0] = 1
        dp[1] = 1
        for i in range(1, len(s)):
            if s[i] == '0':
                if s[i - 1] in ('1', '2'):
                    dp[i + 1] = dp[i - 1]
                else:
                    return 0
            elif s[i - 1] == '1' or s[i - 1] == '2' and '1' <= s[i] <= '6':
                # s[i]单独放置 + s[i-2]单独放置（因为此时s[i]和s[i-1]可以组合）
                dp[i + 1] = dp[i] + dp[i - 1]
            else:
                dp[i + 1] = dp[i]
        return dp[-1]

    def numDecodings(self, s: str) -> int:
        if len(s) == 0 or s[0] == '0':
            return 0
        pre = 1
        cur = 1
        for i in range(1, len(s)):
            if s[i] == '0':
                if s[i - 1] in ('1', '2'):
                    cur, pre = pre, cur
                else:
                    return 0
            elif s[i - 1] == '1' or s[i - 1] == '2' and '1' <= s[i] <= '6':
                cur, pre = cur + pre, cur
            else:
                pre = cur
        return cur


def test():
    sl = Solution()
    for s, target in [
        ("10", 1),
        ("1010", 1),
        ("1213", 5),
        ("12", 2),
        ("1", 1),
        ("0", 0),
        ("226", 3),
        ("28", 1),
        ("0", 0),
    ]:
        ans = sl.numDecodings(s)
        assert ans == target, f"s {s} expect {target}, got {ans}"


if __name__ == "__main__":
    test()
