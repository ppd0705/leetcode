# https://leetcode-cn.com/contest/weekly-contest-199/problems/shuffle-string/

from typing import List


class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        splits = [""] * len(s)
        for i, idx in enumerate(indices):
            splits[idx] = s[i]
        return "".join(splits)


def test():
    solution = Solution()
    for s, indices, target in [
        ("codeleet", [4, 5, 6, 7, 0, 2, 1, 3], "leetcode"),
        ("abc", [0, 1, 2], "abc"),
    ]:
        ans = solution.restoreString(s, indices)
        assert ans == target, f"got {ans}, expect {target}"


if __name__ == "__main__":
    test()
