from typing import List


class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        trie = {}
        end_word = "#"
        tmp_word = "@"
        for word in words:
            children = trie
            for c in word:
                children = children.setdefault(c, {})
            children[end_word] = end_word

        row = len(board)
        column = len(board[0])
        res = []

        def dfs(i: int, j: int, node: dict, s: str):
            v = board[i][j]
            if v not in node:
                return
            s += v
            node = node[v]

            if end_word in node:
                res.append(s)
                node.pop(end_word)  # can only find once

            board[i][j] = tmp_word
            for (dx, dy) in ((0, 1), (0, -1), (1, 0), (-1, 0)):
                x, y = i + dx, j + dy
                if 0 <= x < row and 0 <= y < column and board[x][y] != tmp_word:
                    dfs(x, y, node, s)
            board[i][j] = v

        for r in range(row):
            for c in range(column):
                dfs(r, c, trie, "")

        return res


def test():
    words = ["oath", "pea", "eat", "rain", "ea"]
    expect = ["ea", "eat", "oath"]
    board = [
        ['o', 'a', 'a', 'n'],
        ['e', 't', 'a', 'e'],
        ['i', 'h', 'k', 'r'],
        ['i', 'f', 'l', 'v']
    ]
    ans = Solution().findWords(board, words)
    assert sorted(ans) == expect, f"expect: {expect}, got: {ans}"


if __name__ == "__main__":
    test()

