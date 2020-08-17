from typing import List

class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        f1 = f2 = 0
        for x in cost:
            f2, f1 = min(f2, f1) + x, f2
        return min(f2, f1)
