# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        def helper(node):
            if node is None:
                return 0
            return 1 + max(helper(node.left), helper(node.right))

        return helper(root)
