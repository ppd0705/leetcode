class Solution:
    def firstUniqChar(self, s: str) -> int:
        arr = [0] * 26

        for i, c in enumerate(s):
            arr[ord(c) - 97] = i

        for i, c in enumerate(s):
            if arr[ord(c) - 97] == i:
                return i
            else:
                arr[ord(c) - 97] = -1
        return -1
