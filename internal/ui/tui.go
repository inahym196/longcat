package ui

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/inahym196/longcat/pkg/game"
)

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
			switch cell {
			case game.CellEmpty:
				b.WriteString(".")
			case game.CellFilled:
				b.WriteString("o")
			case game.CellWall:
				b.WriteString("#")
			default:
				panic("invalid cell")
			}
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

		switch ch {
		case 'q':
			return nil
		case 'k':
			c.game.Move(game.DirectionUp)
		case 'j':
			c.game.Move(game.DirectionDown)
		case 'h':
			c.game.Move(game.DirectionLeft)
		case 'l':
			c.game.Move(game.DirectionRight)
		case 'u':
			// not implemented yet
		case 'r':
			// not implemented yet
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
