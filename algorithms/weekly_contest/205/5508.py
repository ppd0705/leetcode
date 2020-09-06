# https://leetcode-cn.com/contest/weekly-contest-205/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/

from typing import List


class Solution:
    def numTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        nums1_count = {}
        nums2_count = {}

        nums1_product_count = {}
        nums2_product_count = {}

        for i in nums1:
            if i in nums1_count:
                nums1_count[i] += 1
            else:
                nums1_count[i] = 1

        for i in nums2:
            if i in nums2_count:
                nums2_count[i] += 1
            else:
                nums2_count[i] = 1

        unique_nums1 = list(nums1_count)
        unique_nums2 = list(nums2_count)
        for i, n in enumerate(unique_nums1):
            if nums1_count[n] > 1:
                nums1_product_count[n * n] = nums1_count[n] * (nums1_count[n] - 1) // 2

            for j in range(i + 1, len(unique_nums1)):
                nums1_product_count[n * unique_nums1[j]] = nums1_count[n] * nums1_count[unique_nums1[j]]

        for i, n in enumerate(unique_nums2):

            if nums2_count[n] > 1:
                nums2_product_count[n * n] = nums2_count[n] * (nums2_count[n] - 1) // 2

            for j in range(i + 1, len(unique_nums2)):
                nums2_product_count[n * unique_nums2[j]] = nums2_count[n] * nums2_count[unique_nums2[j]]
        print({k * k: v for k, v in nums1_count.items()})
        print(nums2_product_count)
        print({k * k: v for k, v in nums2_count.items()})
        print(nums1_product_count)
        ret = 0
        for n in unique_nums1:
            ret += nums2_product_count.get(n * n, 0) * nums1_count[n]
        for n in unique_nums2:
            ret += nums1_product_count.get(n * n, 0) * nums2_count[n]

        return ret


def test():
    sl = Solution()

    for num1, num2, target in [
        #  todo: failed case
        ([7, 1, 7, 3, 6, 2, 6, 4, 2, 6], [5, 5, 2, 2, 7, 1, 7, 2, 6, 7, 6], 30),

        ([4, 7, 9, 11, 23], [3, 5, 1024, 12, 18], 0),
        ([1, 1], [1, 1, 1], 9),
        ([7, 7, 8, 3], [1, 2, 9, 7], 2),
        ([7, 4], [5, 2, 8, 9], 1),
        ([], [], 0),
        ([1], [], 0),
        ([1], [1], 0),
        ([1], [1, 4], 0),
        ([2], [1, 4], 1),
    ]:
        ans = sl.numTriplets(num1, num2)
        assert ans == target, f"ans: {ans}, target: {target}"


if __name__ == "__main__":
    test()
