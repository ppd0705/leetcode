from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 2:
            return 0

        min_price = prices[0]
        ans = 0

        for p in prices[1:]:
            if p > min_price:
                ans = max(ans, p - min_price)
            else:
                min_price = p
        return ans
