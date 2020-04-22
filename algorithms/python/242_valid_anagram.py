class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        letter_count_map_s = {}
        letter_count_map_t = {}

        for i, j in enumerate(s):
            if j in letter_count_map_s:
                letter_count_map_s[j] += 1
            else:
                letter_count_map_s[j] = 1

            k = t[i]
            if k in letter_count_map_t:
                letter_count_map_t[k] += 1
            else:
                letter_count_map_t[k] = 1

        return letter_count_map_s == letter_count_map_t

