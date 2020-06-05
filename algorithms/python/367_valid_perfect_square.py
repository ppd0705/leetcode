class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        i, j = 0, num

        while i <= j:
            mid = (i + j) // 2
            square = mid * mid

            if square == num:
                return True
            elif square > num:
                j = mid - 1
            else:
                i = mid + 1
        return False


def tests():
    s = Solution()
    assert s.isPerfectSquare(1)
    assert s.isPerfectSquare(4)
    assert s.isPerfectSquare(9)
    assert s.isPerfectSquare(16)
    assert not s.isPerfectSquare(2)
    assert not s.isPerfectSquare(3)
    assert not s.isPerfectSquare(5)
    assert not s.isPerfectSquare(6)
    assert not s.isPerfectSquare(7)
    assert not s.isPerfectSquare(8)
    assert not s.isPerfectSquare(10)
    assert not s.isPerfectSquare(11)
    assert not s.isPerfectSquare(15)


if __name__ == "__main__":
    tests()
