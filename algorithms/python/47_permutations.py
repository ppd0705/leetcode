from typing import List


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        def dfs(idx):
            if idx == len(nums):
                ans.append(option[:])
                return

            for i, v in enumerate(visted):
                if v == 0:
                    if i > 0 and nums[i] == nums[i - 1] and not visted[i - 1]:
                        print(idx, i, visted, option[:i])
                        continue

                    option[idx] = nums[i]
                    visted[i] = 1
                    dfs(idx + 1)
                    visted[i] = 0

        ans = []
        nums.sort()
        visted = [0] * len(nums)
        option = [0] * len(nums)
        dfs(0)
        return ans

    def permuteUnique2(self, nums: List[int]) -> List[List[int]]:
        def dfs(idx):
            if idx == len(nums):
                ans.append(option[:])
                return
            for k, v in elements.items():
                if v > 0:
                    option[idx] = k
                    elements[k] -= 1
                    dfs(idx + 1)
                    elements[k] += 1
        ans = []
        option = [0] * len(nums)
        elements = {}

        for n in nums:
            if n in elements:
                elements[n] += 1
            else:
                elements[n] = 1
        dfs(0)
        return ans


def test():
    ans = Solution().permuteUnique2([1, 1, 1, 3])
    print(ans)
    # expect = [[1, 2, 3], [1, 3, 2], [2, 1


if __name__ == "__main__":
    test()
