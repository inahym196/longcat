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
	board *Board
	head  Point
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

func (g *Game) Head() Point { return g.head }

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
	}
	return moved
}
