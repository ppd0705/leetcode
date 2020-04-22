class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        unique_s = set(s)

        for i in unique_s:
            if s.count(i) != t.count(i):
                return False
        return True
