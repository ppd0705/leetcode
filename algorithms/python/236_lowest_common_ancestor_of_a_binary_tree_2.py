class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        parent = {root.val: None}
        visted = {}

        def dfs(node):
            if not node:
                return
            if node.left:
                parent[node.left.val] = node
                dfs(node.left)
            if node.right:
                parent[node.right.val] = node
                dfs(node.right)

        dfs(root)
        while p:
            if p is q:
                return p
            visted[p.val] = 1
            p = parent[p.val]
        while q:
            if q.val in visted:
                return q
            q = parent[q.val]
        return root


def test():
    node1 = TreeNode(1)
    node2 = TreeNode(2)
    node3 = TreeNode(3)
    node4 = TreeNode(4)
    node5 = TreeNode(5)
    node6 = TreeNode(6)

    node1.left = node2
    node1.right = node3
    node2.left = node4
    node2.right = node5
    node3.left = node6

    s = Solution()
    assert s.lowestCommonAncestor(node1, node1, node6) is node1
    assert s.lowestCommonAncestor(node1, node2, node6) is node1
    assert s.lowestCommonAncestor(node1, node3, node6) is node3
    assert s.lowestCommonAncestor(node1, node2, node5) is node2
    assert s.lowestCommonAncestor(node1, node4, node5) is node2


if __name__ == "__main__":
    test()
