class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        def dfs(node: TreeNode):
            if node is not None:
                node.left, node.right = node.right, node.left
                dfs(node.right)
                dfs(node.left)

        dfs(root)
        return root
