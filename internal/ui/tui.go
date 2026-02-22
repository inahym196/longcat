package ui

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/inahym196/longcat/pkg/game"
)

// TODO: uiが増えたタイミングでCommandをusecaseに移す
type Command uint8

const (
	CmdNoop Command = iota
	CmdQuit
	CmdMoveUp
	CmdMoveDown
	CmdMoveLeft
	CmdMoveRight
	CmdUndo
	CmdReset
)

type Input struct {
	Cmd Command
}

type MoveResult struct {
	Moved   bool
	Cleared bool
	Ended   bool
}

type Point struct {
	X, Y int
}

func ParseKey(r rune) Command {
	switch r {
	case 'q':
		return CmdQuit
	case 'k':
		return CmdMoveUp
	case 'j':
		return CmdMoveDown
	case 'h':
		return CmdMoveLeft
	case 'l':
		return CmdMoveRight
	case 'u':
		return CmdUndo
	case 'r':
		return CmdReset
	default:
		return CmdNoop
	}
}

func cellToString(cell game.Cell) string {
	switch cell {
	case game.CellEmpty:
		return "."
	case game.CellFilled:
		return "o"
	case game.CellWall:
		return "#"
	default:
		panic("invalid cell")
	}
}

func Render(g *game.Game) string {
	var b strings.Builder
	for y, row := range g.Board().Cells {
		if y > 0 {
			b.WriteByte('\n')
		}
		for x, cell := range row {
			if x > 0 {
				b.WriteByte(' ')
			}
			p := game.Point{X: x, Y: y}
			if g.Head() == p {
				if cell != game.CellFilled {
					panic(fmt.Sprintf("invalid head cell at (%d,%d)", x, y))
				}
				b.WriteString("H")
				continue
			}
			b.WriteString(cellToString(cell))
		}
	}
	return b.String()
}

type Controller struct {
	in   *bufio.Reader
	out  io.Writer
	game *game.Game
}

func NewController(r io.Reader, w io.Writer, g *game.Game) *Controller {
	return &Controller{bufio.NewReader(r), w, g}
}

func (c *Controller) Run() error {
	for {
		fmt.Fprintln(c.out, Render(c.game))
		ch, _, err := c.in.ReadRune()
		if err != nil {
			return err
		}

		switch ParseKey(ch) {
		case CmdQuit:
			return nil
		case CmdMoveUp:
			c.game.Move(game.DirectionUp)
		case CmdMoveDown:
			c.game.Move(game.DirectionDown)
		case CmdMoveLeft:
			c.game.Move(game.DirectionLeft)
		case CmdMoveRight:
			c.game.Move(game.DirectionRight)
		default:
			// noop
		}
		if c.game.IsCleared() {
			fmt.Fprintln(c.out, Render(c.game))
			fmt.Fprintf(c.out, "CLEARED")
			return nil
		}
	}
}
