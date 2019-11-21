class Solution:
    def isValid(self, s: str) -> bool:
        mapping = {
            "]": "[",
            ")": "(",
            "}": "{",
        }
        stack = []
        for i in s:
            if i in mapping:
                if not stack or stack[-1] != mapping[i]:
                    return False
                else:
                    stack.pop()
            else:
                stack.append(i)

        return len(stack) == 0
