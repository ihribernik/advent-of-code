package directions

import "fmt"

type direction struct {
	x int
	y int
}

var (
	NORTH = NewDirection(0, 1)
	SOUTH = NewDirection(0, -1)
	EAST  = NewDirection(1, 0)
	WEST  = NewDirection(-1, 0)

	DIRECTIONS = map[string]Direction{
		"^": NORTH,
		"v": SOUTH,
		"<": EAST,
		">": WEST,
	}
)

func NewDirection(x, y int) Direction {
	return direction{x: x, y: y}
}

func (d direction) X() int {
	return d.x
}

func (d direction) Y() int {
	return d.y
}

func (d direction) LessThan(dir Direction) bool {
	if d.x != dir.X() {
		return d.x < dir.X()
	}
	return d.y < dir.Y()
}

func (d direction) GreaterThan(dir Direction) bool {
	if d.x != dir.X() {
		return d.x > dir.X()
	}
	return d.y > dir.Y()
}

func (d direction) NewPositionWith(dir Direction) Direction {
	return d.NewPosition(dir, 1)
}

func (d direction) NewPosition(dir Direction, n int) Direction {
	dx := d.x + (n * dir.X())
	dy := d.y + (n * dir.Y())
	return NewDirection(dx, dy)
}

func (d direction) String() string {
	return fmt.Sprintf("%d,%d", d.x, d.y)
}
