from pathlib import Path
from utils import Runner
from typing import List


class Day02(Runner):
    def __init__(self, input_file):
        self.input_file = input_file
        self.multiple_lines = True
        self.input_stream = self.parse(self.input_file)

    def parse_line(self, line) -> List[int]:
        numbers = [int(i) for i in line.strip().split("x")]
        return numbers

    def part1(self) -> str:
        result = 0
        for line in self.input_stream:
            l, w, h = self.parse_line(line)
            areas = (l * w, w * h, h * l)
            total_area = sum(areas) * 2
            extra_paper = min(areas)
            result_area = total_area + extra_paper
            result += result_area

        return str(result)

    def part2(self) -> str:
        result = 0
        for line in self.input_stream:
            numbers = self.parse_line(line)
            l, w, h = numbers
            min_1, min_2 = sorted(numbers)[0:2]
            wrap_feet = min_1 + min_1 + min_2 + min_2
            bow_feet = l * w * h
            total_feet = wrap_feet + bow_feet
            result += total_feet
        return str(result)
