class Solution:
    def longestValidParentheses0(self, s: str) -> int:
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

    def longestValidParentheses(self, s: str) -> int:
        left = right = max_length = 0

        for c in s:
            if c == '(':
                left += 1
            else:
                right += 1

            if left == right:
                max_length = max(max_length, 2 * left)
            elif right > left:
                left = right = 0

        left = right = 0
        for i in range(len(s) - 1, -1, -1):
            c = s[i]
            if c == '(':
                left += 1
            else:
                right += 1

            if left == right:
                max_length = max(max_length, 2 * left)
            elif right < left:
                left = right = 0
        return max_length


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
