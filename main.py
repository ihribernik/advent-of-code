from aoc2015 import day01
from pathlib import Path

ROOT_FOLDER = Path(__name__).cwd()
INPUTS_FOLDER = ROOT_FOLDER / "inputs"
YEARS_FOLDER = {
    "aoc2015": INPUTS_FOLDER / "aoc2015",
}


def main():
    runner = day01.Day01(input_file=YEARS_FOLDER["aoc2015"] / "day01.txt")
    runner.run()


if __name__ == "__main__":
    main()
