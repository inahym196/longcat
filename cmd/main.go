package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	ui "github.com/inahym196/longcat/internal/ui"
	"github.com/inahym196/longcat/pkg/game"
)

func main() {
	rows := []string{
		"########",
		"#H.....#",
		"#......#",
		"########",
	}
	g, err := game.NewGameFromText(rows)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create game: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "controls: h/j/k/l to move, q to quit")
	c := ui.NewController(os.Stdin, os.Stdout, g)
	if err := c.Run(); err != nil && !errors.Is(err, io.EOF) {
		fmt.Fprintf(os.Stderr, "controller error: %v\n", err)
		os.Exit(1)
	}
}
