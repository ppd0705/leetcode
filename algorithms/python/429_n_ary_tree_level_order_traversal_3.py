from typing import List


# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        def traverse_node(node, level):
            if len(ret) == level:
                ret.append([])
            ret[level].append(node.val)

            for child in node.children:
                traverse_node(child, level+1)

        if root is None:
            return []
        ret = []
        traverse_node(root, 0)
        return ret
