package common

type Direction struct {
	X int
	Y int
}

var (
	NORTH = Direction{0, 1}
	SOUTH = Direction{0, -1}
	EAST  = Direction{1, 0}
	WEST  = Direction{-1, 0}
)

func (dir Direction) LessThan(otherDir Direction) bool {
	if dir.X != otherDir.X {
		return dir.X < otherDir.X
	}
	return dir.Y < otherDir.Y
}

func (dir Direction) GreaterThan(otherDir Direction) bool {
	if dir.X != otherDir.X {
		return dir.X > otherDir.X
	}
	return dir.Y > otherDir.Y
}

func (dir Direction) NewPosition(otherDir Direction, n int) Direction {
	dx := dir.X + n + otherDir.X
	dy := dir.Y + n + otherDir.Y
	return Direction{dx, dy}
}
