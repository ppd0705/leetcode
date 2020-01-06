from typing import List


class Solution:
    def minMutation(self, start: str, end: str, bank: List[str]) -> int:
        pass

        bank = set(bank)
        if end not in bank:
            return -1

        change_map = {
            "A": "CGT",
            "C": "AGT",
            "G": "CAT",
            "T": "CGA",
        }
        queue = [(start, 0)]

        while queue:
            print(queue)
            node, step = queue.pop(0)

            if node == end:
                return step

            for i, s in enumerate(node):
                for c in change_map[s]:
                    new = node[:i] + c + node[i+1:]
                    if new in bank:
                        queue.append((new, step+1))
                        bank.remove(new)
        return -1


s = "AAAACCCC"
e = "CCCCCCCC"
b = ["AAAACCCA","AAACCCCA","AACCCCCA","AACCCCCC","ACCCCCCC","CCCCCCCC","AAACCCCC","AACCCCCC"]

# print(Solution().minMutation("AACCGGTT", "AAACGGTA", ["AACCGGTA", "AACCGCTA", "AAACGGTA"]))
print(Solution().minMutation(s, e, b))
# print(Solution().minMutation("AACCGGTT", "AACCGGTA", ["AACCGGT"]))
