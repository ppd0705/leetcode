from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        def generate_parentheses(left, right, string):
            if left == right == 0:
                solutions.append(string)
                return

            if left:
                generate_parentheses(left - 1, right, string + "(")

            if right > left:
                generate_parentheses(left, right - 1, string + ")")

        solutions = []
        generate_parentheses(n, n, "")
        return solutions
