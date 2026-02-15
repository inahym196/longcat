package game_test

import (
	"reflect"
	"testing"

	"github.com/inahym196/longcat/pkg/game"
)

const (
	E = game.CellEmpty
	W = game.CellWall
	F = game.CellFilled
)

func newGameFromASCII(rows []string) *game.Game {
	height := len(rows)
	if height == 0 {
		panic("invalid rows: height == 0")
	}

	var head game.Point
	width := len(rows[0])
	cells := make([][]game.Cell, 0, height)
	for y, row := range rows {
		if len(row) != width {
			panic("rows must be rectangular")
		}
		cellsRow := make([]game.Cell, 0, width)
		for x, ch := range row {
			switch ch {
			case '.':
				cellsRow = append(cellsRow, game.CellEmpty)
			case 'o':
				cellsRow = append(cellsRow, game.CellFilled)
			case 'H':
				cellsRow = append(cellsRow, game.CellFilled)
				head = game.Point{X: x, Y: y}
			case '#':
				cellsRow = append(cellsRow, game.CellWall)
			default:
				panic("invalid cell")
			}
		}
		cells = append(cells, cellsRow)
	}
	return &game.Game{
		Board: &game.Board{
			Width:  width,
			Height: height,
			Cells:  cells,
		},
		Head: head,
	}
}

func TestNewGameFromASCII(t *testing.T) {
	g := newGameFromASCII([]string{".oH#"})

	want := &game.Game{
		Board: &game.Board{
			Width:  4,
			Height: 1,
			Cells:  [][]game.Cell{{E, F, F, W}},
		},
		Head: game.Point{X: 2, Y: 0},
	}

	if !reflect.DeepEqual(g, want) {
		t.Errorf("want %v, got %v", want, g)
	}
}

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	g := newGameFromASCII([]string{"H..#"})
	want := newGameFromASCII([]string{"ooH#"})

	moved := g.Move(game.DirectionRight)

	if !moved {
		t.Fatal("expected move")
	}

	if !reflect.DeepEqual(g, want) {
		t.Errorf("want %v, got %v", want, g)
	}
}

func TestGame_Move_Right_BlockedByWall(t *testing.T) {
	g := newGameFromASCII([]string{"H#"})

	moved := g.Move(game.DirectionRight)

	if moved {
		t.Errorf("expected no move")
	}
}
