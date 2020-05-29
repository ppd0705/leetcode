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

    def generateParenthesis2(self, n: int) -> List[str]:
        solutions = []
        if n == 0:
            return solutions

        queue = [("", n, n)]  # val, left, right

        while queue:
            node = queue.pop(0)
            if node[2] == 0:
                solutions.append(node[0])
                continue

            if node[1] > 0:
                queue.append((node[0] + "(", node[1] - 1, node[2]))
            if node[2] > node[1]:
                queue.append((node[0] + ")", node[1], node[2] - 1))

        return solutions
