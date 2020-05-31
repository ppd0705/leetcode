from typing import List


def has_one_diff(x: str, y: str) -> bool:
    count = 0
    for i in range(len(x)):
        if x[i] != y[i]:
            count += 1
        if count > 1:
            return False
    return count == 1


class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        if endWord not in wordList:
            return 0

        queue: List[str] = [beginWord]
        used = [0] * len(wordList)
        step = 0
        while queue:
            l = len(queue)
            step += 1
            for i in range(l):
                for j, w in enumerate(wordList):
                    if has_one_diff(queue[i], w):
                        if w == endWord:
                            return step + 1
                        if used[j]:
                            continue

                        queue.append(w)
                        used[j] = 1
            queue = queue[l:]
        return 0

    def ladderLength2(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        if endWord not in wordList:
            return 0

        start_queue: List[str] = [beginWord]
        end_queue: List[str] = [endWord]
        used = [0] * len(wordList)
        step = 0
        while start_queue:
            l = len(start_queue)
            step += 1
            for i in range(l):
                for e in end_queue:
                    if has_one_diff(start_queue[i], e):
                        return step + 1
                for j, w in enumerate(wordList):
                    if has_one_diff(start_queue[i], w):

                        if used[j]:
                            continue
                        start_queue.append(w)
                        used[j] = 1
            start_queue = start_queue[l:]
            if len(start_queue) > len(end_queue):
                start_queue, end_queue = end_queue, start_queue
        return 0

    def ladderLength3(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        if endWord not in wordList:
            return 0

        start_queue: List[str] = [beginWord]
        end_queue: List[str] = [endWord]
        start_used = [0] * len(wordList)
        end_used = [0] * len(wordList)
        end_used[wordList.index(endWord)] = 1
        step = 0
        while start_queue:
            l = len(start_queue)
            step += 1
            for i in range(l):
                for j, w in enumerate(wordList):
                    if has_one_diff(start_queue[i], w):
                        if start_used[j]:
                            continue
                        if end_used[j]:
                            return step + 1
                        start_queue.append(w)
                        start_used[j] = 1
            start_queue = start_queue[l:]
            if len(start_queue) > len(end_queue):
                start_queue, end_queue = end_queue, start_queue
                start_used, end_used = end_used, start_used
        return 0

    def ladderLength4(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        start_queue: List[str] = [beginWord]
        end_queue: List[str] = [endWord]
        start_used = [0] * len(wordList)
        end_used = [0] * len(wordList)
        word_map = {w: i for i, w in enumerate(wordList)}
        if endWord not in word_map:
            return 0
        end_used[word_map[endWord]] = 1
        step = 0
        letters = [chr(o) for o in range(ord('a'), ord('z') + 1)]
        while start_queue:
            l = len(start_queue)
            step += 1
            for i in range(l):
                q = start_queue[i]
                q_char = list(q)
                for j in range(len(q)):
                    o = q_char[j]
                    for c in letters:
                        q_char[j] = c
                        s = ''.join(q_char)
                        if s in word_map:
                            if start_used[word_map[s]]:
                                continue
                            if end_used[word_map[s]]:
                                return step + 1
                            start_queue.append(s)
                            start_used[word_map[s]] = 1
                    q_char[j] = o
            start_queue = start_queue[l:]
            if len(start_queue) > len(end_queue):
                start_queue, end_queue = end_queue, start_queue
                start_used, end_used = end_used, start_used
        return 0


def test():
    from collections import namedtuple
    s = Solution()
    sample = namedtuple("test", ["beginWord", "endWord", "wordList", "expect"])
    tests = [
        sample("ymain", "oecij",
               ["ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"], 10),
        sample("a", "c", ["a", "b", "c"], 2),
        sample("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"], 5),
        sample("hit", "cog", ["hot", "dot", "dog", "lot", "log"], 0),
        sample("hit", "cog", ["hot", "cog"], 0),
    ]
    for t in tests:
        ans = s.ladderLength4(t.beginWord, t.endWord, t.wordList)
        assert ans == t.expect, f"failed. expect {t.expect}, got {ans}"


if __name__ == "__main__":
    test()
