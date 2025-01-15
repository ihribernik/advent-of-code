import itertools

from utils import Runner
from utils.direction import Direction

DIRECTIONS = {
    '^': Direction.NORTH,
    'v': Direction.SOUTH,
    '<': Direction.EAST,
    '>': Direction.WEST,
}


class Solution(Runner):
    def __init__(self, input_file):
        self.input_file = input_file
        self.input_stream = self.parse(self.input_file)

    def part1(self) -> int:
        pos = (0, 0)
        already_used = {pos}
        for direction in self.input_stream:
            pos = DIRECTIONS[direction].new_position(*pos)
            already_used.add(pos)

        return len(already_used)

    def part2(self) -> int:
        positions = [(0, 0)] * 2
        already_used = set(positions)
        for op in itertools.batched(self.input_stream, len(positions)):
            for direction, position in zip(op, positions):
                print(direction, position)
                positions = DIRECTIONS[direction].new_position(*position)
                already_used.add(positions)
        return len(already_used)
