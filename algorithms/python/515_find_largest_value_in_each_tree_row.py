from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        ans = []
        if root is None:
            return ans

        queue = [root]
        while queue:
            l = len(queue)
            val = queue[0].val
            for i in range(l):
                if queue[i].val > val:
                    val = queue[i].val

                if queue[i].left:
                    queue.append(queue[i].left)
                if queue[i].right:
                    queue.append(queue[i].right)
            queue = queue[l:]
            ans.append(val)
        return ans
