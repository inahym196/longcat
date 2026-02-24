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

func mustNewCellsAndHeadFromText(t *testing.T, rows []string) ([][]game.Cell, game.Point) {
	cells := make([][]game.Cell, 0, len(rows))

	head := game.Point{}
	for y, row := range rows {
		cellRow := make([]game.Cell, 0, len(row))
		for x, ch := range row {
			switch ch {
			case '.':
				cellRow = append(cellRow, game.CellEmpty)
			case 'o':
				cellRow = append(cellRow, game.CellFilled)
			case 'H':
				cellRow = append(cellRow, game.CellFilled)
				head = game.Point{x, y}
			case '#':
				cellRow = append(cellRow, game.CellWall)
			default:
				t.Fatalf("invalid cell: %q", ch)
			}
		}
		cells = append(cells, cellRow)
	}
	return cells, head
}

func mustNewGameFromText(t *testing.T, rows []string) *game.Game {
	cells, head := mustNewCellsAndHeadFromText(t, rows)
	g, err := game.NewGame(cells, head)
	if err != nil {
		t.Fatal(err)
	}
	return g
}

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	g := mustNewGameFromText(t, []string{"#####", "#H..#", "#####"})
	want := mustNewGameFromText(t, []string{"#####", "#ooH#", "#####"})

	moved := g.Move(game.DirectionRight)
	if !moved {
		t.Fatal("expected move")
	}

	if got := g.Board(); !reflect.DeepEqual(got.Snapshot(), want.Board().Snapshot()) {
		t.Errorf("want %v, got %v", want.Board().Snapshot(), got.Snapshot())
	}
	if got := g.Head(); got != want.Head() {
		t.Errorf("want %v, got %v", want.Head(), got)
	}
}

func TestGame_Move_Right_BlockedByWall(t *testing.T) {
	g := mustNewGameFromText(t, []string{"###", "#H#", "###"})

	moved := g.Move(game.DirectionRight)

	if moved {
		t.Errorf("expected no move")
	}
}

func TestGame_Move_Right_BlockedByFilled(t *testing.T) {
	g := mustNewGameFromText(t, []string{
		"####",
		"#Ho#",
		"####",
	})

	moved := g.Move(game.DirectionRight)

	if moved {
		t.Errorf("expected no move")
	}
}

func TestGame_Move_Directions(t *testing.T) {
	tests := []struct {
		name string
		g    *game.Game
		dir  game.Direction
		want game.Point
	}{
		{"up", mustNewGameFromText(t, []string{"###", "#.#", "#H#", "###"}), game.DirectionUp, game.Point{1, 1}},
		{"down", mustNewGameFromText(t, []string{"###", "#H#", "#.#", "###"}), game.DirectionDown, game.Point{1, 2}},
		{"left", mustNewGameFromText(t, []string{"####", "#.H#", "####"}), game.DirectionLeft, game.Point{1, 1}},
		{"right", mustNewGameFromText(t, []string{"####", "#H.#", "####"}), game.DirectionRight, game.Point{2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Move(tt.dir)
			got := tt.g.Head()

			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestGame_IsCleared_True_NoEmptyCells(t *testing.T) {
	g := mustNewGameFromText(t, []string{"###", "#H#", "###"})

	if !g.IsCleared() {
		t.Errorf("want cleared")
	}
}

func TestGame_IsNotCleared_False_EmptyCellsExists(t *testing.T) {
	g := mustNewGameFromText(t, []string{"####", "#H.#", "####"})

	if g.IsCleared() {
		t.Errorf("want not cleared")
	}
}
