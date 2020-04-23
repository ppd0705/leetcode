from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        from collections import defaultdict
        ret = defaultdict(list)

        for s in strs:
            ret[tuple(sorted(s))].append(s)

        return list(ret.values())
