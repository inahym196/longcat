package ui

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/inahym196/longcat/pkg/game"
)

var (
	TUIParseErr = fmt.Errorf("parse error")
)

func ParseText(rows []string) ([][]game.Cell, game.Point, error) {
	if len(rows) == 0 {
		return nil, game.Point{}, TUIParseErr
	}

	width := len(rows[0])

	cells := make([][]game.Cell, 0, len(rows))
	head := game.Point{}

	for y, row := range rows {

		cellRow := make([]game.Cell, 0, width)
		for x, ch := range row {
			switch ch {
			case '.':
				cellRow = append(cellRow, game.CellEmpty)
			case 'H':
				cellRow = append(cellRow, game.CellEmpty)
				head = game.Point{X: x, Y: y}
			case '#':
				cellRow = append(cellRow, game.CellWall)
			default:
				return nil, game.Point{}, TUIParseErr
			}
		}
		cells = append(cells, cellRow)
	}
	return cells, head, nil
}

func Render(g *game.Game) string {
	var b strings.Builder
	for y, row := range g.Cells() {
		if y > 0 {
			b.WriteByte('\n')
		}
		for x, cell := range row {
			if x > 0 {
				b.WriteByte(' ')
			}
			p := game.Point{X: x, Y: y}
			switch cell {
			case game.CellEmpty:
				if g.Cat().Head() == p {
					b.WriteString("H")
					continue
				}
				if g.Cat().IsBody(p) {
					b.WriteString("o")
					continue
				}
				b.WriteString(".")
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
