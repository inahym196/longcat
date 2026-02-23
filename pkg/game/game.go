package game

import (
	"fmt"
	"slices"
)

type Game struct {
	board      *Board
	head       Point
	emptyCount int
}

var (
	GameNoWallErr = fmt.Errorf("壁が必要")
)

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

	return NewGame(&Board{
		Width:  width,
		Height: len(rows),
		Cells:  cells,
	}, head)
}

func NewGame(b *Board, h Point) (*Game, error) {
	if b == nil {
		return nil, fmt.Errorf("board is nil")
	}
	if slices.IndexFunc(b.Cells[0], func(c Cell) bool { return c != CellWall }) != -1 {
		return nil, GameNoWallErr
	}
	if slices.IndexFunc(b.Cells[len(b.Cells)-1], func(c Cell) bool { return c != CellWall }) != -1 {
		return nil, GameNoWallErr
	}
	for _, row := range b.Cells {
		if row[0] != CellWall {
			return nil, GameNoWallErr
		}
		if row[len(row)-1] != CellWall {
			return nil, GameNoWallErr
		}
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
		if !g.board.InBounds(next) {
			panic("something wrong")
		}
		if g.board.IsWall(next) {
			break
		}
		if g.board.IsFilled(next) {
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
