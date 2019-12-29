from typing import List

number_letters_map = {
    '2': ['a', 'b', 'c'],
    '3': ['d', 'e', 'f'],
    '4': ['g', 'h', 'i'],
    '5': ['j', 'k', 'l'],
    '6': ['m', 'n', 'o'],
    '7': ['p', 'q', 'r', 's'],
    '8': ['t', 'u', 'v'],
    '9': ['w', 'x', 'y', 'z'],
}


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        ret = [""]

        for s in digits:
            next_ret = []
            for r in ret:
                for letter in number_letters_map[s]:
                    next_ret.append(r + letter)
            ret = next_ret

        return ret


print(Solution().letterCombinations("234"))
