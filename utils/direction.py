from enum import Enum
from typing import Tuple


class Direction(Enum):
    NORTH = (0, 1)
    SOUTH = (0, -1)
    EAST = (1, 0)
    WEST = (-1, 0)

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __lt__(self, other: object) -> bool:
        if not isinstance(other, Direction):
            return NotImplemented
        else:
            return (self.x, self.y) < (other.x, other.y)

    def __gt__(self, other: object) -> bool:
        if not isinstance(other, Direction):
            return NotImplemented
        else:
            return (self.x, self.y) > (other.x, other.y)

    def new_position(self, x, y, *, n=1) -> Tuple[int, int]:
        return self.x + n * x, self.y + n * y
