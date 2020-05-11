from typing import List
from collections import deque


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
        ret = []
        curr = root
        while curr is not None:
            if curr.left is None:
                ret.append(curr.val)
                curr = curr.right
            else:
                pre = curr.left
                while pre.right is not None:
                    pre = pre.right

                pre.right = curr

                curr = curr.left
                pre.right.left = None
        return ret
