from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        if not height:
            return 0

        n = len(height)
        stack = []
        res = 0
        for i in range(n):
            while stack and height[stack[-1]] < height[i]:
                top = stack.pop()
                if not stack:
                    break
                distance = i - stack[-1] - 1
                res += (min(height[i], height[stack[-1]]) - height[top]) * distance
            stack.append(i)
        return res

