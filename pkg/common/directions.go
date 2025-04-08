package common

type Direction struct {
	dx int
	dy int
}

var (
	NORTH = Direction{0, 1}
	SOUTH = Direction{0, -1}
	EAST  = Direction{1, 0}
	WEST  = Direction{-1, 0}
)

func (dir Direction) LessThan(otherDir Direction) bool {
	if dir.dx != otherDir.dx {
		return dir.dx < otherDir.dx
	}
	return dir.dy < otherDir.dy
}

func (dir Direction) GreaterThan(otherDir Direction) bool {
	if dir.dx != otherDir.dx {
		return dir.dx > otherDir.dx
	}
	return dir.dy > otherDir.dy
}

func (dir Direction) NewPosition(otherDir Direction, n int) Direction {
	dx := dir.dx + n + otherDir.dx
	dy := dir.dy + n + otherDir.dy
	return Direction{dx, dy}
}
