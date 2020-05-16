class Solution:
    def climbStairs(self, n: int) -> int:

        cache = {}

        def climb_stairs(n):
            if n in cache:
                return cache[n]
            if n <= 2:
                ret = n
            else:
                ret = climb_stairs(n - 1) + climb_stairs(n - 2)
            cache[n] = ret
            return ret

        return climb_stairs(n)
