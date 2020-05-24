from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        def dfs(idx):
            if idx == len(nums):
                ans.append(option[:])
                return

            for i, v in enumerate(visted):
                if v == 0:
                    option[idx] = nums[i]
                    visted[i] = 1
                    dfs(idx + 1)
                    visted[i] = 0

        ans = []
        visted = [0] * len(nums)
        option = [0] * len(nums)
        dfs(0)
        return ans

    def permute2(self, nums: List[int]) -> List[List[int]]:
        def dfs(idx):
            if idx == len(nums):
                ans.append(option[:])
                return
            for i in range(idx, len(nums)):
                option[idx] = nums[i]
                nums[i], nums[idx] = nums[idx], nums[i]
                dfs(idx + 1)
                nums[i], nums[idx] = nums[idx], nums[i]

        ans = []
        option = [0] * len(nums)
        dfs(0)
        return ans


def test():
    ans = Solution().permute([1, 2, 3])
    expect = [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
    assert ans == expect, ans
    print(Solution().permute2([1, 2, 3]))


if __name__ == "__main__":
    test()
