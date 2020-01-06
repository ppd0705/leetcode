from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        def dfs(node, level):
            if node is None:
                return
            if len(result) < level + 1:
                result.append([])
            result[level].append(node.val)

            for child in (node.left, node.right):
                dfs(child, level + 1)

        result = []
        dfs(root, 0)
        return result


