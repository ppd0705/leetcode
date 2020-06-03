from typing import List


class Solution:
    def robotSim(self, commands: List[int], obstacles: List[List[int]]) -> int:
        pointer = [0, 0]
        idx = 1
        sign = 1

        command_map = {
            # (idx, sign, command) -> (idx, sign)
            (1, 1, -2): (0, -1),
            (1, 1, -1): (0, 1),
            (1, -1, -2): (0, 1),
            (1, -1, -1): (0, -1),
            (0, 1, -2): (1, 1),
            (0, 1, -1): (1, -1),
            (0, -1, -2): (1, -1),
            (0, -1, -1): (1, 1),
        }
        obstacles_set = set((o[0], o[1]) for o in obstacles)
        max_distance = 0
        for c in commands:
            if c == -2 or c == -1:
                idx, sign = command_map[(idx, sign, c)]
            else:
                for i in range(c):
                    pointer[idx] += 1 * sign
                    if (pointer[0], pointer[1]) in obstacles_set:
                        pointer[idx] -= 1 * sign
                        break
                distance = pointer[0] * pointer[0] + pointer[1] * pointer[1]
                if distance > max_distance:
                    max_distance = distance
        return max_distance

    def robotSim2(self, commands: List[int], obstacles: List[List[int]]) -> int:
        x, y, dx, dy = 0, 0, 0, 1
        obstacles_set = set((o[0], o[1]) for o in obstacles)
        max_distance = 0
        for c in commands:
            if c == -2:
                dx, dy = -dy, dx
            elif c == -1:
                dx, dy = dy, -dx
            else:
                for i in range(c):
                    next_x = x + dx
                    next_y = y + dy
                    if (next_x, next_y) in obstacles_set:
                        break
                    x = next_x
                    y = next_y
                distance = x ** 2 + y ** 2
                if distance > max_distance:
                    max_distance = distance
        return max_distance


def test():
    from collections import namedtuple
    sample = namedtuple("test", ["commands", "obstacles", "expect"])

    tests = [
        sample([4, -1, 3], [], 25),
        sample([4, -1, 4, -2, 4], [[2, 4]], 65),
    ]
    s = Solution()
    for t in tests:
        ans = s.robotSim2(t.commands, t.obstacles)
        assert ans == t.expect, f"expect {t.expect}, got {ans}"


if __name__ == "__main__":
    test()
