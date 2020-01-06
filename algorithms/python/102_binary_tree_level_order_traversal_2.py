from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        queue = [(0, root)]
        result = []
        while queue:
            level, node = queue.pop(0)
            if node is None:
                continue
            if len(result) < level + 1:
                result.append([])
            result[-1].append(node.val)

            for child in (node.left, node.right):
                queue.append((level+1, child))

        return result