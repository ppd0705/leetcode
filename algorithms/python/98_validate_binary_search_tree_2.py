class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        stack, current = [], float("-inf")
        while stack or root:
            while root is not None:
                stack.append(root)
                root = root.left
            root = stack.pop()

            if root.val <= current:
                return False
            current = root.val
            root = root.right

        return True
