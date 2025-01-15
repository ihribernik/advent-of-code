from utils import Runner


class Solution(Runner):
    def __init__(self, input_file: str) -> None:
        self.input_file = input_file
        self.input_stream = self.parse(input_file)

    def part1(self) -> int:
        result = 0
        for direction in self.input_stream:
            match direction:
                case "(":
                    result += 1
                case ")":
                    result -= 1
        return result

    def part2(self) -> int:
        floor = 0
        position_char = 0
        for index, direction in enumerate(self.input_stream):
            match direction:
                case "(":
                    floor += 1
                case ")":
                    floor -= 1
                    if floor == -1:
                        position_char = index + 1
                        break

        return position_char
