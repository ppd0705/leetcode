class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        counter = [0] * 26
        ord_a = ord("a")
        for i, j in enumerate(s):
            counter[ord(j) - ord_a] += 1
            counter[ord(t[i]) - ord_a] -= 1

        for c in counter:
            if c != 0:
                return False
        return True




