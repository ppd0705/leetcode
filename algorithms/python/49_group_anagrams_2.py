from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        from collections import defaultdict
        ret = defaultdict(list)
        prime = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103]

        for s in strs:
            key = 1
            for k in s:
                key *= prime[ord(k) - 97]
            ret[key].append(s)

        return list(ret.values())
