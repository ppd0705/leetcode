from typing import List


class Solution:
    def maxProfit(self, prices: List[int], fee: int) -> int:
        hold, not_hold = -prices[0] - fee, 0

        for p in prices[1:]:
            hold, not_hold = max(hold, not_hold - p - fee), max(not_hold, hold + p)
        return not_hold
