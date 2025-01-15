from abc import ABC, abstractmethod
from logging import Logger

logger = Logger(__name__)


class Runner(ABC):
    input_file: str
    multiple_lines: bool = False

    @abstractmethod
    def __init__(self, input_file: str):
        raise NotImplementedError("Runner.__init__ is not Implemented")

    @property
    def name(self) -> str:
        return self.input_file

    @staticmethod
    def parse(file_name: str) -> str:
        with open(file_name, "r", encoding="utf-8") as fr:
            return fr.read()

    @abstractmethod
    def part1(self) -> int:
        raise NotImplementedError("Runner.part1 is not Implemented")

    @abstractmethod
    def part2(self) -> int:
        raise NotImplementedError("Runner.part2 is not Implemented")

    def run(self):
        result_1 = self.part1()
        result_2 = self.part2()
        logger.warning("current day: %s", self.name)
        logger.warning("part 1: %s", result_1)
        logger.warning("part 2: %s", result_2)
