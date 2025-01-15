import pytest

from aoc2015.day03 import Solution
from tests.fixtures.day03 import input_stream, second_input_stream


@pytest.mark.parametrize(
    ('input_s', 'expected'),
    (
        *input_stream,
    ),
)
def test_day03_part_1(input_s: str, expected: int, monkeypatch):
    day = Solution("inputs/aoc2015/day03.txt")
    monkeypatch.setattr(day, "input_stream", input_s)
    result = day.part1()
    assert result == expected


@pytest.mark.parametrize(
    ('input_s', 'expected'),
    (
        *second_input_stream,
    ),
)
def test_day03_part_2(input_s: str, expected: int, monkeypatch):
    day = Solution("inputs/aoc2015/day03.txt")
    monkeypatch.setattr(day, "input_stream", input_s)
    result = day.part2()
    assert result == expected
