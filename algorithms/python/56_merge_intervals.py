from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])
        ret = []
        for i in intervals:
            if not ret or ret[-1][1] < i[0]:
                ret.append(i)
            else:
                ret[-1][1] = max(ret[-1][1], i[1])
        return ret


def test():
    s = Solution()
    for (intervals, target) in [
        ([[1, 3], [2, 6], [8, 10], [15, 18]], [[1, 6], [8, 10], [15, 18]]),
        ([[1, 4], [4, 5]], [[1, 5]]),
    ]:
        ans = s.merge(intervals)
        assert ans == target, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
