class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        def dfs(node, limit_down, limit_up) -> bool:
            if node is None:
                return True
            val = node.val
            return limit_down < val < limit_up and dfs(node.left, limit_down, val) and dfs(node.right, val, limit_up)

        return dfs(root, float("-inf"), float("inf"))
