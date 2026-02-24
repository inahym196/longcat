package game_test

import (
	"slices"
	"testing"

	"github.com/inahym196/longcat/pkg/game"
)

const (
	E = game.CellEmpty
	W = game.CellWall
)

func mustNewGameFromText(t *testing.T, rows []string) *game.Game {
	stage := mustNewStageFromText(t, rows)
	g, err := game.NewGame(stage.Board(), stage.Head())
	if err != nil {
		t.Fatal(err)
	}
	return g
}

func mustNewStageFromText(t *testing.T, rows []string) game.Stage {
	cells := make([][]game.Cell, 0, len(rows))
	var head game.Point
	for y, row := range rows {
		cellRow := make([]game.Cell, 0, len(row))
		for x, ch := range row {
			switch ch {
			case '.':
				cellRow = append(cellRow, game.CellEmpty)
			case '#':
				cellRow = append(cellRow, game.CellWall)
			case 'H':
				cellRow = append(cellRow, game.CellEmpty)
				head = game.Point{x, y}
			default:
				t.Fatalf("invalid cell: %q", ch)
			}
		}
		cells = append(cells, cellRow)
	}
	stage, err := game.NewStage(cells, head)
	if err != nil {
		t.Fatal(err)
	}
	return stage
}

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	g := mustNewGameFromText(t, []string{"#####", "#H..#", "#####"})
	wantBody := []game.Point{{1, 1}, {2, 1}}
	wantHead := game.Point{3, 1}

	moved := g.Move(game.DirectionRight)
	if !moved {
		t.Fatal("expected move")
	}
	got := g.Cat()

	if got.Head() != wantHead {
		t.Errorf("want %v, got %v", wantHead, got.Head())
	}
	if !slices.Equal(got.Body(), wantBody) {
		t.Errorf("want %v, got %v", wantBody, got.Body())
	}
}

func TestGame_Move_Right_BlockedByWall(t *testing.T) {
	g := mustNewGameFromText(t, []string{"###", "#H#", "###"})

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
			got := tt.g.Cat().Head()

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
