class Solution:
    def hammingWeight(self, n: int) -> int:
        ans = 0
        for s in bin(n)[2:]:
            ans += s == '1'
        return ans

    def hammingWeight1(self, n: int) -> int:
        return bin(n).count('1')

    def hammingWeight2(self, n: int) -> int:
        ans = 0
        while n:
            n, m = divmod(n, 2)
            ans += m == 1
        return ans

    def hammingWeight3(self, n: int) -> int:
        ans = 0
        while n:
            ans += n & 1
            n >>= 1
        return ans

    def hammingWeight4(self, n: int) -> int:
        ans = 0
        while n:
            ans += 1
            n &= n-1
        return ans


def test():
    s = Solution()
    for (n, target) in ((0, 0), (1, 1), (10, 2), (2**32-1, 32)):
        ans = s.hammingWeight3(n)
        assert ans == target, f"target {target}, got {ans}"


if __name__ == "__main__":
    test()
