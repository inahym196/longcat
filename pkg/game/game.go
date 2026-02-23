package game

import (
	"fmt"
)

type Game struct {
	board      *Board
	head       Point
	emptyCount int
}

func NewGameFromText(rows []string) (*Game, error) {
	if len(rows) == 0 {
		return nil, fmt.Errorf("rows is empty")
	}

	width := len(rows[0])
	if width == 0 {
		return nil, fmt.Errorf("rows[0] is empty")
	}

	cells := make([][]Cell, 0, len(rows))
	head := Point{}
	headFound := false

	for y, row := range rows {
		if len(row) != width {
			return nil, fmt.Errorf("rows must be rectangular")
		}

		cellRow := make([]Cell, 0, width)
		for x, ch := range row {
			switch ch {
			case '.':
				cellRow = append(cellRow, CellEmpty)
			case 'o':
				cellRow = append(cellRow, CellFilled)
			case 'H':
				if headFound {
					return nil, fmt.Errorf("head must be single")
				}
				cellRow = append(cellRow, CellFilled)
				head = Point{X: x, Y: y}
				headFound = true
			case '#':
				cellRow = append(cellRow, CellWall)
			default:
				return nil, fmt.Errorf("invalid cell: %q", ch)
			}
		}
		cells = append(cells, cellRow)
	}
	if !headFound {
		return nil, fmt.Errorf("head not found")
	}
	b, err := NewBoard(width, len(rows), cells)
	if err != nil {
		panic(err)
	}
	return NewGame(b, head)
}

func NewGame(b *Board, h Point) (*Game, error) {
	if b == nil {
		return nil, fmt.Errorf("board is nil")
	}
	return &Game{b, h, b.EmptyCount()}, nil
}

func (g *Game) Head() Point { return g.head }

func (g *Game) Board() Board {
	cells := make([][]Cell, len(g.board.Cells))
	for y, row := range g.board.Cells {
		cells[y] = append([]Cell(nil), row...)
	}
	return Board{
		Width:  g.board.Width,
		Height: g.board.Height,
		Cells:  cells,
	}
}

func (g *Game) Move(d Direction) bool {

	current := g.head
	moved := false

	for {
		next := current.Move(d)
		nextCell, err := g.board.Cell(next)
		if err != nil {
			panic(err)
		}
		if nextCell != CellEmpty {
			break
		}
		g.board.Fill(next)
		current = next
		moved = true
	}

	if moved {
		g.head = current
		g.emptyCount = g.board.EmptyCount()
	}
	return moved
}

func (g *Game) IsCleared() bool {
	return g.emptyCount == 0
}

// TODO: Game.EmptyCountは必要になったら公開する
