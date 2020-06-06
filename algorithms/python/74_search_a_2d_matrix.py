from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        row = len(matrix)
        if row == 0:
            return False
        column = len(matrix[0])
        if column == 0:
            return False

        for i in range(row):
            if matrix[i][column - 1] >= target:
                l, r = 0, column - 1
                while l <= r:
                    mid = (l + r) // 2
                    if matrix[i][mid] == target:
                        return True
                    elif matrix[i][mid] < target:
                        l = mid + 1
                    else:
                        r = mid - 1

                return False
        return False


def test():
    s = Solution()
    assert s.searchMatrix([[1, 3]], 3)


if __name__ == "__main__":
    test()
