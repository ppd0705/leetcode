from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = [[-1, 0], ]  # pair: (index, height)
        max_area = 0
        for i, h in enumerate(heights):
            if h > stack[-1][1]:
                stack.append([i, h])
            elif h == stack[-1][1]:
                stack[-1][0] = i
            else:
                while h < stack[-1][1]:
                    max_area = max(max_area, stack[-1][1] * (i - stack[-2][0] - 1))
                    stack.pop()
                stack.append([i, h])

        for i in range(1, len(stack)):
            max_area = max(max_area, stack[i][1] * (stack[-1][0] - stack[i - 1][0]))

        return max_area
