package game

import (
	"fmt"
	"slices"
)

type Cell int

const (
	CellEmpty Cell = iota
	CellWall
)

var (
	BoardInvalidHeightErr = fmt.Errorf("height is invalid")
	BoardInvalidWidthErr  = fmt.Errorf("width is invalid")
	BoardNoTopWallErr     = fmt.Errorf("no top wall")
	BoardNoBottomWallErr  = fmt.Errorf("no bottom wall")
	BoardNoSideWallErr    = fmt.Errorf("no side wall")
)

type Board struct {
	width      int
	height     int
	cells      [][]Cell
	emptyCount int
}

func NewBoard(cells [][]Cell) (Board, error) {
	if len(cells) < 3 {
		return Board{}, BoardInvalidHeightErr
	}
	if len(cells[0]) < 3 {
		return Board{}, BoardInvalidWidthErr
	}
	b := Board{len(cells[0]), len(cells), cells, countEmpty(cells)}

	if err := b.validateBoard(); err != nil {
		return Board{}, err
	}
	return b, nil
}

func countEmpty(cells [][]Cell) int {
	count := 0
	for _, rows := range cells {
		for _, c := range rows {
			if c == CellEmpty {
				count++
			}
		}
	}
	return count
}

func (b Board) EmptyCount() int { return b.emptyCount }

func (b Board) validateBoard() error {
	if slices.IndexFunc(b.cells[0], func(c Cell) bool { return c != CellWall }) != -1 {
		return BoardNoTopWallErr
	}
	if slices.IndexFunc(b.cells[b.height-1], func(c Cell) bool { return c != CellWall }) != -1 {
		return BoardNoBottomWallErr
	}
	for _, row := range b.cells {
		if row[0] != CellWall {
			return BoardNoSideWallErr
		}
		if row[len(row)-1] != CellWall {
			return BoardNoSideWallErr
		}
	}
	return nil
}

func (b Board) IsInside(p Point) bool {
	return 0 <= p.Y && p.Y < len(b.cells) && 0 <= p.X && p.X < len(b.cells[0])
}

func (b Board) IsWall(p Point) bool { return b.cells[p.Y][p.X] == CellWall }

func (b Board) Cells() [][]Cell { return slices.Clone(b.cells) }
