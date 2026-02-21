package ui

import (
	"testing"

	"github.com/inahym196/longcat/pkg/game"
)

const (
	E = game.CellEmpty
	W = game.CellWall
	F = game.CellFilled
)

func TestRender(t *testing.T) {
	ss := Snapshot{
		Width:  5,
		Height: 3,
		Cells: [][]game.Cell{
			{W, W, W, W, W},
			{W, F, E, F, W},
			{W, W, W, W, W},
		},
		Head: Point{X: 3, Y: 1},
	}

	got := Render(ss)
	want := "[# # # # #\n# o . H #\n# # # # #]"
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestRender_HeadMustBeFilled(t *testing.T) {
	ss := Snapshot{
		Width:  3,
		Height: 3,
		Cells: [][]game.Cell{
			{W, W, W},
			{W, E, W},
			{W, W, W},
		},
		Head: Point{X: 1, Y: 1},
	}

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()

	_ = Render(ss)
}
