# https://leetcode-cn.com/contest/weekly-contest-205/problems/replace-all-s-to-avoid-consecutive-repeating-characters/


class Solution:
    def modifyString(self, s: str) -> str:
        str_list = list(s)
        for i, c in enumerate(str_list):
            if c != "?":
                continue
            for j in range(97, 97 + 26):
                if i - 1 >= 0 and str_list[i - 1] == chr(j):
                    continue
                if i + 1 < len(str_list) and str_list[i + 1] == chr(j):
                    continue
                str_list[i] = chr(j)
                break
        return "".join(str_list)


def test():
    sl = Solution()
    for s, target in [
        ("???", "aba"),
        ("a?", "ab"),
        ("?a?", "bab"),
        ("", ""),
        ("?", "a"),
        ("abc", "abc"),
    ]:
        ans = sl.modifyString(s)
        assert ans == target, f"s: {s}, target: {target}, ans: {ans}"


if __name__ == "__main__":
    test()
