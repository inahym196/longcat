package game

import "fmt"

type Game struct {
	board      *Board
	head       Point
	emptyCount int
}

var (
	GameHeadErr = fmt.Errorf("invalid head")
)

func NewGame(cells [][]Cell, h Point) (*Game, error) {

	b, err := NewBoard(cells)
	if err != nil {
		return nil, err
	}

	headCell, err := b.Cell(h)
	if err != nil {
		return nil, err
	}

	if headCell != CellFilled {
		return nil, GameHeadErr
	}

	return &Game{b, h, b.EmptyCount()}, nil
}

func (g *Game) Head() Point { return g.head }

func (g *Game) Board() *Board {
	cells := make([][]Cell, len(g.board.Cells))
	for y, row := range g.board.Cells {
		cells[y] = append([]Cell(nil), row...)
	}
	return &Board{
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
