from typing import List


class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        counter = [0] * 1001

        for i in arr1:
            counter[i] += 1

        idx = 0
        for i in arr2:
            while counter[i] > 0:
                arr1[idx] = i
                idx += 1
                counter[i] -= 1

        for i, c in enumerate(counter):
            for j in range(c):
                arr1[idx] = i
                idx += 1
        return arr1

