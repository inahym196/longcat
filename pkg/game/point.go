package game

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Point struct {
	X, Y int
}

func (p Point) Move(d Direction) Point {
	switch d {
	case DirectionUp:
		return p.moveUp()
	case DirectionDown:
		return p.moveDown()
	case DirectionLeft:
		return p.moveLeft()
	case DirectionRight:
		return p.moveRight()
	default:
		panic("invalid direction")
	}
}

func (p Point) moveUp() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) moveDown() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) moveLeft() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) moveRight() Point {
	return Point{p.X + 1, p.Y}
}
