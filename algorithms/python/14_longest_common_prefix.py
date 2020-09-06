from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""

        ans = strs[0]
        idx = 0

        go_on = True
        while go_on and idx < len(ans):
            for i in range(1, len(strs)):
                if idx >= len(strs[i]) or strs[i][idx] != ans[idx]:
                    go_on = False
                    break
            else:
                idx += 1
        return ans[:idx]


def test():
    sl = Solution()
    for strs, target in [
        (["a", "ab", "b"], ""),
        (["a", "ab"], "a"),
        (["a", "b"], ""),
        (["a"], "a"),
        ([], ""),

    ]:
        ans = sl.longestCommonPrefix(strs)
        assert ans == target, f"ans: {ans}, target: {target}"


if __name__ == "__main__":
    test()
