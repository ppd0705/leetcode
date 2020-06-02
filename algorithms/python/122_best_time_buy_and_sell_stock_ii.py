from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        amount = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i - 1]:
                amount += prices[i] - prices[i - 1]
        return amount
