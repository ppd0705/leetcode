class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        i, j, ret = 0, len(height) - 1,  0
        while i < j:
            if height[i] < height[j]:
                area = height[i] * (j - i)
                i += 1
            else:
                area = height[j] * (j - i)
                j -= 1
            if area > ret:
                ret = area
        return ret
