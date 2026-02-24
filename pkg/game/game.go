package game

type Stage struct {
	board Board
	head  Point
}

func NewStage(cells [][]Cell, head Point) (Stage, error) {
	b, err := NewBoard(cells)
	if err != nil {
		return Stage{}, err
	}
	return Stage{b, head}, nil
}

func (s Stage) Head() Point  { return s.head }
func (s Stage) Board() Board { return s.board }

type Game struct {
	board       Board
	initialHead Point
	cat         *LongCat
}

func NewGame(board Board, head Point) (*Game, error) {
	cat, err := NewLongCat(head, board)
	if err != nil {
		return nil, err
	}
	return &Game{board, head, cat}, nil
}

func (g *Game) IsCleared() bool {
	return g.cat.Len() == g.board.EmptyCount()
}

func (g *Game) Move(d Direction) bool {
	return g.cat.Stretch(d, g.board)
}

func (g *Game) Cells() [][]Cell  { return g.board.Cells() }
func (g *Game) Cat() LongCatView { return g.cat }
