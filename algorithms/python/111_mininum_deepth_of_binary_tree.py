class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if root is None:
            return 0

        queue = [root]
        depth = 0
        while queue:
            depth += 1
            l = len(queue)
            for i in range(l):
                if queue[i].left is None and queue[i].right is None:
                    return depth

                if queue[i].left is not None:
                    queue.append(queue[i].left)
                if queue[i].right is not None:
                    queue.append(queue[i].right)

            queue = queue[l:]
        return depth
