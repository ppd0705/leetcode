from typing import List


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def __init__(self):
        self._traverse_path = []

    def inorderTraversal(self, root: TreeNode) -> List[int]:
        self.inorder(root)
        return self._traverse_path

    def inorder(self, root):
        if root:
            self.inorder(root.left)
            self._traverse_path.append(root.val)
            self.inorder(root.right)