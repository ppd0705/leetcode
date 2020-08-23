from typing import List


class Solution:
    def findLatestStep(self, arr: List[int], m: int) -> int:
        if m == len(arr):
            return len(arr)
        mark = [0] * (len(arr) + 2)
        ans = -1
        # todo: need to record index range
        # index_tuple = []
        for i, a in enumerate(arr):
            mark[a] = 1
            cnt = 1
            for j in range(a - 1, 0, -1):
                if mark[j] == 1:
                    cnt += 1
                else:
                    break
            for j in range(a + 1, len(arr) + 1):
                if mark[j] == 1:
                    cnt += 1
                else:
                    break
            if cnt == m or ans == i:
                ans = i + 1
        return ans


def test():
    s = Solution()
    for arr, m, target in [
        ([3, 5, 1, 2, 4], 1, 4),
        ([1, 2], 1, 1),
    ]:
        ans = s.findLatestStep(arr, m)
        assert ans == m, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
