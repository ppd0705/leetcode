class Solution:
    def mySqrt(self, x: int) -> int:
        i, j = 0, x
        while i < j:
            mid = (i + j + 1) // 2
            square = mid * mid
            if square == x:
                return x
            elif square > x:
                j = mid - 1
            else:
                i = mid
        return i
