from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return

        root = TreeNode(preorder[0])
        mid = inorder.index(preorder[0])

        root.left = self.buildTree(preorder[1:mid + 1], inorder[:mid])
        root.right = self.buildTree(preorder[mid + 1:], inorder[mid + 1:])
        return root


def test():
    preorder = [3, 9, 20, 15, 7]
    inorder = [9, 3, 15, 20, 7]

    node = Solution().buildTree(preorder, inorder)

    assert node.val == 3
    assert node.left.val == 9
    assert node.right.left.val == 15


if __name__ == "__main__":
    test()
