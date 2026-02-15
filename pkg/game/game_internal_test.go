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

func TestGame_Move_Right_MovesUntilWall(t *testing.T) {
	b := game.Board{4, 1, []game.Cell{F, E, E, W}}
	h := game.Point{0, 0}
	g := &game.Game{b, h}

	wantB := game.Board{4, 1, []game.Cell{F, F, F, W}}
	wantH := game.Point{2, 0}
	wantG := &game.Game{wantB, wantH}

	moved := g.Move(game.DirectionRight)

	if !moved {
		t.Fatal("expected move")
	}

	if !reflect.DeepEqual(g, wantG) {
		t.Errorf("want %v, got %v", wantG, g)
	}
}
