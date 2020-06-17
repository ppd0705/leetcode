from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0

        prev_not_hold = 0
        prev_hold = -prices[0]
        curr_not_hold = max(0, prev_hold+prices[1])
        curr_hold = max(prev_not_hold-prices[1], prev_hold)

        for p in prices[2:]:
            prev_hold, curr_hold = curr_hold, max(curr_hold, prev_not_hold - p)
            prev_not_hold, curr_not_hold = curr_not_hold, max(curr_not_hold, prev_hold + p)
        return curr_not_hold
