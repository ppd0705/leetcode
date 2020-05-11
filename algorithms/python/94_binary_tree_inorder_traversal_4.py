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
        ret = []
        WHITE, GRAY = 0, 1
        stack = [(root, WHITE)]
        while stack:
            node, color = stack.pop()
            if node is None:
                continue
            if color == WHITE:
                stack.append((node.right, WHITE))
                stack.append((node, GRAY))
                stack.append((node.left, WHITE))
            else:
                ret.append(node.val)
        return ret


print(Solution().inorderTraversal(None))

n1 = TreeNode(1)
n2 = TreeNode(2)
n3 = TreeNode(3)
n4 = TreeNode(4)
n1.left = n2
n1.right = n3
n3.left = n4
print(Solution().inorderTraversal(n1))
