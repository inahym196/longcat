package ui

import (
	"reflect"
	"testing"

	"github.com/inahym196/longcat/pkg/game"
)

const (
	E = game.CellEmpty
	W = game.CellWall
)

func TestParseText(t *testing.T) {
	rows := []string{
		"#####",
		"#..H#",
		"#####",
	}
	wantCells := [][]game.Cell{
		{W, W, W, W, W},
		{W, E, E, E, W},
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
	cells := [][]game.Cell{
		{W, W, W, W, W},
		{W, E, E, E, W},
		{W, W, W, W, W},
	}
	b, err := game.NewBoard(cells)
	if err != nil {
		t.Fatal(err)
	}
	g, err := game.NewGame(b, game.Point{X: 3, Y: 1})
	if err != nil {
		t.Fatal(err)
	}
	want := "# # # # #\n# . . H #\n# # # # #"

	got := Render(g)

	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
