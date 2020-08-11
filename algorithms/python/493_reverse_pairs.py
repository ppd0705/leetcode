import bisect


class Solution:
    def reversePairs(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = []
        count = 0
        for i in range(len(nums) - 1, -1, -1):
            count += bisect.bisect_left(s, nums[i])
            bisect.insort_left(s, nums[i] * 2)
        return count


def test():
    s = Solution()

    for nums, target in [
        ([5, 4, 3, 2, 1], 4),
    ]:
        ans = s.reversePairs(nums)
        assert ans == target, f"except {target}, got {ans}"


if __name__ == "__main__":
    test()
