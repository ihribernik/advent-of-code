
from pathlib import Path
from utils.importers import modulos

ROOT_FOLDER = Path(__name__).cwd()
INPUTS_FOLDER = ROOT_FOLDER / "inputs"
YEARS_FOLDER = {
    "aoc2015": INPUTS_FOLDER / "aoc2015",
}

def main():
    for name in modulos:
        module = modulos[name]
        solution = module.Solution((YEARS_FOLDER["aoc2015"] / name ).as_posix() + ".txt" )
        solution.run()


if __name__ == "__main__":
    main()
