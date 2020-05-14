from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def __init__(self):
        self._traversal_path = []

    def postorder(self, root: Node) -> List[int]:
        self.post_order(root)
        return self._traversal_path

    def post_order(self, node):
        if node:
            for c in node.children:
                self.post_order(c)
            self._traversal_path.append(node.val)
