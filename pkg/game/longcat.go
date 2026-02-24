package game

import (
	"fmt"
	"slices"
)

var (
	LongCatOutOfBoardErr  = fmt.Errorf("LongCatはボード外にいてはいけない")
	LongCatStuckInWallErr = fmt.Errorf("LongCatは壁の中にいてはいけない")
)

type LongCat struct {
	head Point
	body []Point
}

func NewLongCat(head Point, b Board) (*LongCat, error) {
	if !b.IsInside(head) {
		return nil, LongCatOutOfBoardErr
	}
	if b.IsWall(head) {
		return nil, LongCatStuckInWallErr
	}
	return &LongCat{head, make([]Point, 0)}, nil
}

func (cat *LongCat) Head() Point {
	return cat.head
}

func (cat *LongCat) stretch(next Point) {
	cat.body = append(cat.body, cat.head)
	cat.head = next
}

func (cat *LongCat) Stretch(d Direction, b Board) bool {
	stretched := false
	for {
		next := cat.Head().Move(d)
		if b.IsWall(next) || cat.IsBody(next) {
			break
		}
		cat.stretch(next)
		stretched = true
	}
	return stretched
}

func (cat *LongCat) IsBody(p Point) bool { return slices.Contains(cat.body, p) }
func (cat *LongCat) Len() int            { return len(cat.body) + 1 }
func (cat *LongCat) Body() []Point       { return slices.Clone(cat.body) }

type LongCatView interface {
	Len() int
	Head() Point
	Body() []Point
	IsBody(p Point) bool
}
