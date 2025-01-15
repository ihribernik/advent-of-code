from typing import List

from utils import Runner


class Solution(Runner):
    def __init__(self, input_file: str):
        self.input_file = input_file
        self.multiple_lines = True
        self.input_stream = self.parse(self.input_file)
        if self.multiple_lines:
            self.input_stream = self.input_stream.splitlines()

    @staticmethod
    def parse_line(line: str) -> List[int]:
        numbers = [int(i) for i in line.strip().split("x")]
        return numbers

    def part1(self) -> int:
        result = 0
        for line in self.input_stream:
            l, w, h = self.parse_line(line)
            areas = (l * w, w * h, h * l)
            total_area = sum(areas) * 2
            extra_paper = min(areas)
            result_area = total_area + extra_paper
            result += result_area

        return result

    def part2(self) -> int:
        result = 0
        for line in self.input_stream:
            numbers = self.parse_line(line)
            l, w, h = numbers
            min_1, min_2 = sorted(numbers)[0:2]
            wrap_feet = min_1 + min_1 + min_2 + min_2
            bow_feet = l * w * h
            total_feet = wrap_feet + bow_feet
            result += total_feet
        return result
