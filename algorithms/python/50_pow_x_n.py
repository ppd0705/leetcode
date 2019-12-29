class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        elif n == 1:
            return x

        if n < 0:
            x = 1/x
            n = -n
        sub = self.myPow(x, n // 2)

        if n % 2:
            return x * sub * sub
        else:
            return sub * sub

