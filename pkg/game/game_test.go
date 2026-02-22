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

func mustNewGameFromText(rows []string) *game.Game {
	g, err := game.NewGameFromText(rows)
	if err != nil {
		panic(err)
	}
	return g
}

func TestNewGameFromText(t *testing.T) {
	g, _ := game.NewGameFromText([]string{
		"#####",
		"#.oH#",
		"#####",
	})

	wantB := &game.Board{5, 3, [][]game.Cell{
		{W, W, W, W, W},
		{W, E, F, F, W},
		{W, W, W, W, W},
	}}
	wantP := game.Point{3, 1}

	if !reflect.DeepEqual(g.Board(), wantB) {
		t.Errorf("want %v, got %v", wantB, g.Board())
	}
	if g.Head() != wantP {
		t.Errorf("want %v, got %v", wantP, g.Head())
	}
}

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	g := mustNewGameFromText([]string{
		"#####",
		"#H..#",
		"#####",
	})
	want := mustNewGameFromText([]string{
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
	g := mustNewGameFromText([]string{
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
	g := mustNewGameFromText([]string{
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
		{"up", mustNewGameFromText([]string{"###", "#.#", "#H#", "###"}), game.DirectionUp, game.Point{1, 1}},
		{"down", mustNewGameFromText([]string{"###", "#H#", "#.#", "###"}), game.DirectionDown, game.Point{1, 2}},
		{"left", mustNewGameFromText([]string{"####", "#.H#", "####"}), game.DirectionLeft, game.Point{1, 1}},
		{"right", mustNewGameFromText([]string{"####", "#H.#", "####"}), game.DirectionRight, game.Point{2, 1}},
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
	g := mustNewGameFromText([]string{"###", "#H#", "###"})

	if !g.IsCleared() {
		t.Errorf("want cleared")
	}
}

func TestGame_IsNotCleared_False_EmptyCellsExists(t *testing.T) {
	g := mustNewGameFromText([]string{"####", "#H.#", "####"})

	if g.IsCleared() {
		t.Errorf("want not cleared")
	}
}
