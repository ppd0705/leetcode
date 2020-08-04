# https://leetcode-cn.com/contest/weekly-contest-198/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

from typing import List
import copy


# todo: timeout, need to improve the implement.

class Solution:
    def countSubTrees(self, n: int, edges: List[List[int]], labels: str) -> List[int]:
        ret = [1] + [0] * (n - 1)
        parent_lables = {0: {}}

        for edge in edges:
            a, b = edge
            if b in parent_lables:
                # swap
                a, b = b, a

            ret[b] += 1
            parent_lables[b] = copy.deepcopy(parent_lables[a])
            parent_lables[b].setdefault(labels[a], []).append(a)
            for i in parent_lables[b].get(labels[b], []):
                ret[i] += 1
        return ret


def test():
    s = Solution()
    for (n, edges, lables, target) in [
        (7, [[0, 1], [1, 2], [2, 3], [2, 4], [3, 5], [0, 6]], "cdcecff", [3, 1, 2, 1, 1, 1, 1]),
        (4, [[0, 1], [1, 2], [0, 3]], "bbbb", [4, 2, 1, 1]),
        (7, [[0, 1], [1, 2], [2, 3], [3, 4], [4, 5], [5, 6]], "aaabaaa", [6, 5, 4, 1, 3, 2, 1]),
    ]:
        ans = s.countSubTrees(n, edges, lables)
        assert ans == target, f"ans: {ans}, target: {target}"


if __name__ == "__main__":
    test()
