class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        cnt = 0

        for i in range(len(s)-1, -1, -1):
            if s[i] == ' ':
                if cnt == 0:
                    continue
                else:
                    break
            else:
                cnt += 1
        return cnt
