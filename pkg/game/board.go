package game

import (
	"fmt"
	"slices"
)

type Cell uint8

const (
	CellEmpty Cell = iota
	CellWall
	CellFilled
)

var (
	BoardInvalidHeightErr = fmt.Errorf("height is invalid")
	BoardInvalidWidthErr  = fmt.Errorf("width is invalid")
	BoardNoTopWallErr     = fmt.Errorf("no top wall")
	BoardNoBottomWallErr  = fmt.Errorf("no bottom wall")
	BoardNoSideWallErr    = fmt.Errorf("no side wall")
	BoardOutOfRangeErr    = fmt.Errorf("out of range")
)

type Board struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func NewBoard(cells [][]Cell) (*Board, error) {
	if len(cells) < 3 {
		return nil, BoardInvalidHeightErr
	}
	if len(cells[0]) < 3 {
		return nil, BoardInvalidWidthErr
	}
	b := Board{len(cells[0]), len(cells), cells}

	if err := b.validateBoard(); err != nil {
		return nil, err
	}
	return &b, nil
}

func (b *Board) validateBoard() error {
	if slices.IndexFunc(b.Cells[0], func(c Cell) bool { return c != CellWall }) != -1 {
		return BoardNoTopWallErr
	}
	if slices.IndexFunc(b.Cells[b.Height-1], func(c Cell) bool { return c != CellWall }) != -1 {
		return BoardNoBottomWallErr
	}
	for _, row := range b.Cells {
		if row[0] != CellWall {
			return BoardNoSideWallErr
		}
		if row[len(row)-1] != CellWall {
			return BoardNoSideWallErr
		}
	}
	return nil
}
func (b *Board) Cell(p Point) (Cell, error) {
	if !b.inBounds(p) {
		return CellEmpty, BoardOutOfRangeErr
	}
	return b.Cells[p.Y][p.X], nil
}

func (b *Board) inBounds(p Point) bool {
	return 0 <= p.Y && p.Y < len(b.Cells) && 0 <= p.X && p.X < len(b.Cells[0])
}

func (b *Board) Fill(p Point) {
	b.Cells[p.Y][p.X] = CellFilled
}

func (b *Board) EmptyCount() int {
	count := 0
	for _, rows := range b.Cells {
		for _, c := range rows {
			if c == CellEmpty {
				count++
			}
		}
	}
	return count
}

func (b *Board) Snapshot() [][]Cell {
	return slices.Clone(b.Cells)
}
