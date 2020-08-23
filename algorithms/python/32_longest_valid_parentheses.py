class Solution:
    def longestValidParentheses(self, s: str) -> int:
        if not s:
            return 0
        dp = [0] * len(s)
        for i in range(1, len(s)):
            if s[i] == ')':
                if s[i - 1] == '(':
                    dp[i] = (dp[i - 2] if i >= 2 else 0) + 2
                elif i - dp[i - 1] > 0 and s[i - dp[i - 1] - 1] == '(':
                    dp[i] = dp[i - 1] + 2 + (0 if i - dp[i - 1] - 2 < 0 else dp[i - dp[i - 1] - 2])
        return max(dp)


def test():
    sl = Solution()
    for s, target in [
        ("()", 2),
        ("())", 2),
        (")(())", 4),
        (")()())", 4),
    ]:
        ans = sl.longestValidParentheses(s)
        assert ans == target, f"target: {target}, ans {ans}"


if __name__ == "__main__":
    test()
