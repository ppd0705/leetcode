from typing import List
from typing import Dict

from collections import namedtuple


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


Point = namedtuple("Point", "row col")


class Solution:
    def verticalTraversal(self, root: TreeNode) -> List[List[int]]:
        point_vals_map: Dict[Point, List[int]] = {}

        def dfs(node: TreeNode, row: int, col: int):
            if node is None:
                return
            point = Point(row, col)
            if point not in point_vals_map:
                point_vals_map[point] = []

            point_vals_map[point].append(node.val)
            dfs(node.left, row + 1, col - 1)
            dfs(node.right, row + 1, col + 1)

        dfs(root, 0, 0)

        sorted_items = sorted(point_vals_map.items(), key=lambda x: (x[0].col, x[0].row))

        res = []

        current_col = -10002

        for point, vals in sorted_items:
            vals.sort()
            if point.col != current_col:
                res.append(vals[:])
                current_col = point.col
            else:
                res[-1].extend(vals)

        return res

    def test(self):
        node1 = TreeNode(1)
        node2 = TreeNode(2)
        node3 = TreeNode(3)
        node4 = TreeNode(4)
        node5 = TreeNode(5)
        node6 = TreeNode(6)
        node7 = TreeNode(7)

        assert self.verticalTraversal(node1) == [[1]], f"got {self.verticalTraversal(node1)}"

        node1.left = node2
        assert self.verticalTraversal(node1) == [[2], [1]], f"got {self.verticalTraversal(node1)}"

        node1.right = node3
        node2.left = node4
        node2.right = node7
        node3.left = node5
        node5.left = node6
        assert self.verticalTraversal(node1) == [[4], [2, 6], [1, 5, 7], [3]], f"got {self.verticalTraversal(node1)}"


if __name__ == "__main__":
    Solution().test()
