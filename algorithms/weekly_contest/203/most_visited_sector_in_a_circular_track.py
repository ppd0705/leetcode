from typing import List


# https://leetcode-cn.com/contest/weekly-contest-203/problems/most-visited-sector-in-a-circular-track/
class Solution:
    def mostVisited(self, n: int, rounds: List[int]) -> List[int]:
        arr = [0] * (n + 1)
        for i in range(0, len(rounds) - 1):
            if rounds[i] <= rounds[i + 1]:
                for j in range(rounds[i], rounds[i + 1] + 1):
                    arr[j] += 1
            else:
                for j in range(rounds[i], n + 1):
                    arr[j] += 1

                for j in range(1, rounds[i + 1] + 1):
                    arr[j] += 1
            rounds[i + 1] = ((rounds[i + 1] + 1) % (n + 1)) or 1
        max_cnt = max(arr[1:])
        return [i for i, v in enumerate(arr) if i != 0 and v == max_cnt]


def test():
    s = Solution()
    for n, rounds, target in [
        (2, [2, 1, 2, 1, 2, 1, 2, 1, 2], [2]),
        (4, [1, 3, 1, 2], [1, 2]),
        (7, [1, 3, 5, 7], [1, 2, 3, 4, 5, 6, 7]),
    ]:
        ans = s.mostVisited(n, rounds)
        assert ans == target, f"expect {target}, got {ans}"


if __name__ == "__main__":
    test()
