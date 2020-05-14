from typing import List


# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        if root is None:
            return []
        level = [root]
        ret = []
        while level:
            ret.append([n.val for n in level])
            level = [child for l in level for child in l.children]
        return ret
