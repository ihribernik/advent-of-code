from typing import List

from utils import Runner


class Day03(Runner):
    def __init__(self, input_file):
        self.input_file = input_file
        self.input_stream = self.parse(self.input_file)

    def part1(self) -> int:
        result = 0
        prev_post = ""
        for current_post in self.input_stream:
            if current_post.strip() == prev_post:
                result += 1
                prev_post = current_post
                continue

        return result

    def part2(self) -> int:
        result = 0

        return result
