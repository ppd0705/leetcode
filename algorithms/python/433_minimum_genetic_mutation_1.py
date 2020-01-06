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

        min_count = len(bank) + 1

        def dfs(current, count, current_bank):
            # terminator
            nonlocal min_count

            if count > min_count:
                return
            if current == end:
                if count < min_count:
                    min_count = count
                return

            if not bank:
                return

            # process
            for i, s in enumerate(current):
                for char in change_map[s]:
                    new = current[:i] + char + current[i + 1:]
                    if new not in current_bank:
                        continue
                    current_bank.remove(new)
                    # drill down
                    dfs(new, count + 1, current_bank)

                    # reverse state
                    current_bank.add(new)

        dfs(start, 0, bank)

        return min_count if min_count <= len(bank) else -1
