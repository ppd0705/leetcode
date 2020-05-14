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

    def preorderTraversal(self, root: TreeNode) -> List[int]:
        ret = []
        curr = root
        while curr is not None:
            ret.append(curr.val)
            if curr.left is None:
                curr = curr.right
            else:
                pre = curr.left
                while pre.right is not None:
                    pre = pre.right

                pre.right = curr.right
                curr = curr.left
        return ret
