class DNode:
    __slots__ = ("prev", "next", "key", "value")


class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self._cache = {}
        self._cache_len = self._cache.__len__
        self._root = DNode()
        self._root.prev = self._root.next = self._root

    def _move_to_end(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

        prev = self._root.prev
        prev.next = self._root.prev = node
        node.next = self._root
        node.prev = prev

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self._cache:
            return -1

        node = self._cache[key]
        self._move_to_end(node)
        return node.value

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """

        if key in self._cache:
            node = self._cache[key]
            node.value = value
            self._move_to_end(node)
        else:
            if self._cache_len() < self.capacity:
                prev = self._root.prev
                node = DNode()
                node.prev, node.next, node.key, node.value = prev, self._root, key, value
                prev.next = self._root.prev = node

                self._cache[key] = node
            else:
                node = self._root
                node.key = key
                node.value = value
                self._cache[key] = node

                root = self._root.next
                self._cache.pop(root.key)
                root.key = root.value = None
                self._root = root
