class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = {}
        self.end_word = "#"

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        node = self.root
        for w in word:
            node = node.setdefault(w, {})
        node[self.end_word] = self.end_word

    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        node = self.root
        for w in word:
            node = node.get(w)
            if node is None:
                return False
        return self.end_word in node

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        node = self.root
        for w in prefix:
            node = node.get(w)
            if node is None:
                return False
        return True


def test():
    s = Trie()
    s.insert("apple")
    assert not s.search("app")
    assert s.search("apple")
    assert s.startsWith("app")
    assert s.startsWith("apple")
    s.insert("app")
    assert s.search("app")


if __name__ == "__main__":
    test()
