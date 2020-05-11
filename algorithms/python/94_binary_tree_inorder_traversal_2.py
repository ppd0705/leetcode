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
        stack = deque()
        ret = []
        curr = root
        while curr is not None or len(stack):
            while curr is not None:
                stack.append(curr)
                curr = curr.left
            curr = stack.pop()
            ret.append(curr.val)
            curr = curr.right
        return ret
