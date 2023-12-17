from utils import Runner


class Day01(Runner):
    def __init__(self, input_file):
        self.input_file = input_file
        self.input_stream = self.parse(input_file)

    def name(self) -> str:
        return self.input_file

    def parse(self, file_name):
        with open(file_name, "r", encoding="utf-8") as fr:
            return fr.readline()

    def part1(self) -> str:
        result = 0
        for direction in self.input_stream:
            match direction:
                case "(":
                    result += 1
                case ")":
                    result -= 1

        return str(result)

    def part2(self) -> str:
        floor = 0
        position_char = 0
        for index, direction in enumerate(self.input_stream):
            match direction:
                case "(":
                    floor += 1
                case ")":
                    floor -= 1
                    if floor == -1:
                        position_char = index
                        break

        return str(position_char)

    def run(self):
        return super().run()
