package ui

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

func TestParseText(t *testing.T) {
	rows := []string{
		"#####",
		"#.oH#",
		"#####",
	}
	wantCells := [][]game.Cell{
		{W, W, W, W, W},
		{W, E, F, F, W},
		{W, W, W, W, W},
	}
	wantHead := game.Point{X: 3, Y: 1}

	gotCells, gotHead, err := ParseText(rows)

	if err != nil {
		t.Errorf("want no error, got %v", err)
	}

	if !reflect.DeepEqual(wantCells, gotCells) {
		t.Errorf("want %v, got %v", wantCells, gotCells)
	}
	if wantHead != gotHead {
		t.Errorf("want %v, got %v", wantHead, gotHead)
	}
}

func TestRender(t *testing.T) {
	g, _ := game.NewGame(
		[][]game.Cell{
			{W, W, W, W, W},
			{W, F, E, F, W},
			{W, W, W, W, W},
		},
		game.Point{X: 3, Y: 1},
	)
	want := "# # # # #\n# o . H #\n# # # # #"

	got := Render(g)

	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestRender_HeadMustBeFilled(t *testing.T) {
	g, _ := game.NewGame(
		[][]game.Cell{
			{W, W, W},
			{W, E, W},
			{W, W, W},
		},
		game.Point{X: 1, Y: 1},
	)

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()

	Render(g)
}
