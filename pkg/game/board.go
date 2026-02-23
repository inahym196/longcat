package game

type Cell uint8

const (
	CellEmpty Cell = iota
	CellWall
	CellFilled
)

type Board struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func (b *Board) InBounds(p Point) bool {
	return 0 <= p.Y && p.Y < len(b.Cells) && 0 <= p.X && p.X < len(b.Cells[0])
}

func (b *Board) IsWall(p Point) bool {
	return b.Cells[p.Y][p.X] == CellWall
}

func (b *Board) IsFilled(p Point) bool {
	return b.Cells[p.Y][p.X] == CellFilled
}

func (b *Board) Fill(p Point) {
	b.Cells[p.Y][p.X] = CellFilled
}

func (b *Board) EmptyCount() int {
	count := 0
	for _, rows := range b.Cells {
		for _, c := range rows {
			if c == CellEmpty {
				count++
			}
		}
	}
	return count
}
