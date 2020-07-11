from typing import List
from collections import deque


class Solution:
    def slidingPuzzle(self, board: List[List[int]]) -> int:
        neighbors = [(1, 3), (0, 2, 4), (1, 5), (0, 4), (1, 3, 5), (2, 4)]

        current = tuple(i for j in board for i in j)
        zero_index = current.index(0)
        target = (1, 2, 3, 4, 5, 0)
        queue = deque([(current, zero_index, 0)])  # state zero_index, step
        visited = set(current)

        while queue:
            state, zero_index, step = queue.popleft()
            if state == target:
                return step
            state_list = list(state)
            for n in neighbors[zero_index]:
                state_list[n], state_list[zero_index] = state_list[zero_index], state_list[n]
                new_state = tuple(state_list)
                if new_state not in visited:
                    visited.add(new_state)
                    queue.append((new_state, n, step + 1))
                state_list[n], state_list[zero_index] = state_list[zero_index], state_list[n]
        return -1
