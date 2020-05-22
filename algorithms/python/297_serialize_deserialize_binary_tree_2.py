class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


from collections import deque


class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        res = []
        queue = deque([root])
        while queue:
            node = queue.popleft()
            if node:
                queue.append(node.left)
                queue.append(node.right)
                res.append(str(node.val))
            else:
                res.append("#")

        return ",".join(res)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        parts = data.split(",")

        idx = 0
        val = parts[idx]

        if val == "#":
            return None

        root = TreeNode(int(val))
        queue = deque([root])
        while queue:
            node = queue.popleft()
            idx += 1
            val = parts[idx]
            if val != "#":
                node.left = left = TreeNode(int(val))
                queue.append(left)
            idx += 1
            val = parts[idx]
            if val != "#":
                node.right = right = TreeNode(int(val))
                queue.append(right)
        return root


def test():
    c = Codec()

    n1 = TreeNode(1)
    n2 = TreeNode(2)
    n3 = TreeNode(3)
    n4 = TreeNode(4)
    n5 = TreeNode(5)
    n1.left = n2
    n1.right = n3
    n2.right = n4
    n4.left = n5

    for node in (None, n1, n2, n3, n4, n5):
        data = c.serialize(node)
        n = c.deserialize(data)
        r = c.serialize(n)
        assert r == data, f"{node} {data} {n} {r}"


if __name__ == "__main__":
    test()
