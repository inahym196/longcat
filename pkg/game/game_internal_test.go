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
	g, err := game.NewGame(&game.Board{
		Width:  width,
		Height: height,
		Cells:  cells,
	}, head)
	if err != nil {
		panic(err)
	}
	return g
}

func TestNewGameFromASCII(t *testing.T) {
	g := newGameFromASCII([]string{
		"#####",
		"#.oH#",
		"#####",
	})

	want, err := game.NewGame(&game.Board{5, 3, [][]game.Cell{
		{W, W, W, W, W},
		{W, E, F, F, W},
		{W, W, W, W, W},
	}}, game.Point{3, 1})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(g, want) {
		t.Errorf("want %v, got %v", want, g)
	}
}

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	g := newGameFromASCII([]string{
		"#####",
		"#H..#",
		"#####",
	})
	want := newGameFromASCII([]string{
		"#####",
		"#ooH#",
		"#####",
	})

	moved := g.Move(game.DirectionRight)

	if !moved {
		t.Fatal("expected move")
	}

	if !reflect.DeepEqual(g, want) {
		t.Errorf("want %v, got %v", want, g)
	}
}

func TestGame_Move_Right_BlockedByWall(t *testing.T) {
	g := newGameFromASCII([]string{
		"###",
		"#H#",
		"###",
	})

	moved := g.Move(game.DirectionRight)

	if moved {
		t.Errorf("expected no move")
	}
}

func TestGame_Move_Right_BlockedByFilled(t *testing.T) {
	g := newGameFromASCII([]string{
		"####",
		"#Ho#",
		"####",
	})

	moved := g.Move(game.DirectionRight)

	if moved {
		t.Errorf("expected no move")
	}
}
