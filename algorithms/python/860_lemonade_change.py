from typing import List


class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        m5, m10 = 0, 0

        for b in bills:
            if b == 5:
                m5 += 1
            elif b == 10:
                if m5 > 0:
                    m5 -= 1
                    m10 += 1
                else:
                    return False
            else:
                if m10 > 0 and m5 > 0:
                    m10 -= 1
                    m5 -= 1
                elif m5 > 2:
                    m5 -= 3
                else:
                    return False
        return True
