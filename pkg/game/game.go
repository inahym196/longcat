package game

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Cell uint8

const (
	CellEmpty Cell = iota
	CellWall
	CellFilled
)

type Point struct {
	X, Y int
}

type Board struct {
	Width  int
	Height int
	Cells  []Cell
}

type Game struct {
	Board Board
	Head  Point
}

func (g *Game) Move(d Direction) bool {
	g.Board.Cells = []Cell{CellFilled, CellFilled, CellFilled, CellWall}
	g.Head = Point{2, 0}
	return true
}
