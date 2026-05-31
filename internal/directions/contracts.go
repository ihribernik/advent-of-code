package directions

type Direction interface {
	X() int
	Y() int
	LessThan(dir Direction) bool
	GreaterThan(dir Direction) bool
	NewPosition(dir Direction, n int) Direction
	NewPositionWith(dir Direction) Direction
	String() string
}
