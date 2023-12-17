from abc import ABC, abstractmethod
from pathlib import Path
from typing import List, Tuple
from logging import Logger

logger = Logger(__name__)


class Runner(ABC):
    @abstractmethod
    def __init__(self, input_file):
        raise NotImplementedError("Runner.__init__ is not Implemented")

    @abstractmethod
    def name(self) -> Path:
        raise NotImplementedError("Runner.name is not Implemented")

    @abstractmethod
    def parse(self, file_name):
        raise NotImplementedError("Runner.parse is not Implemented")

    @abstractmethod
    def part1(self) -> str:
        raise NotImplementedError("Runner.part1 is not Implemented")

    @abstractmethod
    def part2(self) -> str:
        raise NotImplementedError("Runner.part2 is not Implemented")

    @abstractmethod
    def run(self):
        result_1 = self.part1()
        result_2 = self.part2()
        logger.warning("current day: %s", self.name().name)
        logger.warning("part 1: %s", result_1)
        logger.warning("part 2: %s", result_2)
