package common

import "fmt"

type Direction struct {
	X int
	Y int
}

var (
	NORTH = Direction{0, 1}
	SOUTH = Direction{0, -1}
	EAST  = Direction{1, 0}
	WEST  = Direction{-1, 0}

	DIRECTIONS = map[string]Direction{
		"^": NORTH,
		"v": SOUTH,
		"<": EAST,
		">": WEST,
	}
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

func (dir Direction) NewPositionWith(otherDir Direction) Direction {
	return dir.NewPosition(otherDir, 1)
}

func (dir Direction) NewPosition(otherDir Direction, n int) Direction {
	dx := dir.X + (n * otherDir.X)
	dy := dir.Y + (n * otherDir.Y)
	return Direction{X: dx, Y: dy}
}

func (d Direction) String() string {
	return fmt.Sprintf("%d,%d", d.X, d.Y)
}
