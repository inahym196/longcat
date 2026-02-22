package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/inahym196/longcat/internal/terminal"
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

	restore, err := terminal.EnterRawMode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to enable raw mode: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := restore(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to restore terminal: %v\n", err)
		}
	}()

	fmt.Fprintln(os.Stdout, "controls: h/j/k/l to move, q to quit")
	c := ui.NewController(os.Stdin, os.Stdout, g)
	if err := c.Run(); err != nil && !errors.Is(err, io.EOF) {
		fmt.Fprintf(os.Stderr, "controller error: %v\n", err)
		os.Exit(1)
	}
}
