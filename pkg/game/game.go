package game

import (
	"fmt"
	"slices"
)

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

func (p Point) MoveRight() Point {
	return Point{p.X + 1, p.Y}
}

type Board struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func (b *Board) InBounds(p Point) bool {
	return 0 <= p.Y && p.Y < len(b.Cells) && 0 <= p.X && p.X < len(b.Cells[0])
}

func (b *Board) IsWall(p Point) bool {
	return b.Cells[p.Y][p.X] == CellWall
}

func (b *Board) IsFilled(p Point) bool {
	return b.Cells[p.Y][p.X] == CellFilled
}

func (b *Board) Fill(p Point) {
	b.Cells[p.Y][p.X] = CellFilled
}

type Game struct {
	Board *Board
	Head  Point
}

var (
	GameNoWallErr = fmt.Errorf("壁が必要")
)

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
	return &Game{b, h}, nil
}

func (g *Game) Move(d Direction) bool {
	if d != DirectionRight {
		panic("not implemented yet")
	}

	current := g.Head
	moved := false

	for {
		next := current.MoveRight()
		if !g.Board.InBounds(next) {
			break
		}
		if g.Board.IsWall(next) {
			break
		}
		if g.Board.IsFilled(next) {
			break
		}
		g.Board.Fill(next)
		current = next
		moved = true
	}

	if moved {
		g.Head = current
	}
	return moved
}
