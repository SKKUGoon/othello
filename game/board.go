package game

import (
	"fmt"
	"strings"
)

type Board struct {
	// Board Dimension
	MySize int
	// Element
	Elem         []*BoardElement
	ElemExistMap map[BoardCoord]bool
	ElemCoord    map[BoardCoord]*BoardElement
	// Game Status
	Whites    int
	Blacks    int
	MyTurn    int // -1 for Black, 1 for White
	TotalTurn int
	// Is game finished
	IsGame bool
}

type Othelo interface {
	// Creating a new board for a new game
	New(int)
	Connect()
	Initialize()

	// Playable check
	possibleMove()

	// Representation of the board
	Picture()
}

func (b *Board) New(size int) {
	// Setup
	b.MySize = size
	b.MyTurn = 1
	b.TotalTurn = 0

	b.Elem = []*BoardElement{}
	b.ElemExistMap = map[BoardCoord]bool{}
	b.ElemCoord = map[BoardCoord]*BoardElement{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			c := BoardCoord{
				X: x,
				Y: y,
			}

			// Single blank space with no stone(0)
			be := BoardElement{
				Coordinate: c,
				Color:      0,
			}

			b.Elem = append(b.Elem, &be)
			b.ElemExistMap[c] = true
			b.ElemCoord[c] = &be
		}
	}
}

func (b *Board) Connect() {
	c := BoardCoord{}
	for _, e := range b.Elem {
		west, east := e.Coordinate.X-1, e.Coordinate.X+1
		south, north := e.Coordinate.Y-1, e.Coordinate.Y+1

		if north < b.MySize {
			// Set North
			c = BoardCoord{X: e.Coordinate.X, Y: north}
			e.N = b.ElemCoord[c]
		}

		if south >= 0 {
			// Set South
			c = BoardCoord{X: e.Coordinate.X, Y: south}
			e.S = b.ElemCoord[c]
		}

		if west >= 0 {
			// Set West
			c = BoardCoord{X: west, Y: e.Coordinate.Y}
			e.W = b.ElemCoord[c]

			if south >= 0 {
				// Set South West
				c = BoardCoord{X: west, Y: south}
				e.SW = b.ElemCoord[c]
			}

			if north < b.MySize {
				// North West
				c = BoardCoord{X: west, Y: north}
				e.NW = b.ElemCoord[c]
			}
		}

		if east < b.MySize {
			// Set East
			c = BoardCoord{X: east, Y: e.Coordinate.Y}
			e.E = b.ElemCoord[c]

			if south >= 0 {
				// Set South East
				c = BoardCoord{X: east, Y: south}
				e.SE = b.ElemCoord[c]
			}

			if north < b.MySize {
				// Set North East
				c = BoardCoord{X: east, Y: north}
				e.NE = b.ElemCoord[c]
			}
		}
	}
}

func (b *Board) Initialize() {
	// Place 4 stones on the middle of the board
	// Stone should be placed diagonally
	centerStart := b.MySize/2 - 1
	centerEnd := centerStart + 1
	// Black
	black1 := BoardCoord{
		X: centerStart,
		Y: centerStart,
	}
	black2 := BoardCoord{
		X: centerEnd,
		Y: centerEnd,
	}
	b.ElemCoord[black1].Color = -1
	b.ElemCoord[black2].Color = -1
	// White
	white1 := BoardCoord{
		X: centerStart,
		Y: centerEnd,
	}
	white2 := BoardCoord{
		X: centerEnd,
		Y: centerStart,
	}
	b.ElemCoord[white1].Color = 1
	b.ElemCoord[white2].Color = 1

	b.Whites = 2
	b.Blacks = 2
}

func (b *Board) isLegalLoc(c BoardCoord) error {
	ok := b.ElemExistMap[c]
	if !ok {
		return &GameError{Code: "G002", Message: "Illegal placement. Outside Dimension"}
	}
	return nil
}

func (b *Board) Move(place BoardCoord, whoseTurn int) error {
	// Placement check `place`
	if err := b.isLegalLoc(place); err != nil {
		return err
	}

	// Turn check
	if b.MyTurn != whoseTurn {
		return &GameError{Code: "G001", Message: "Illegal turn"}
	}

	targetElem := b.ElemCoord[place]
	candidate, err := targetElem.TurnOverCheck(whoseTurn)
	if err != nil {
		return err
	}

	// Update color of the othelo board - where it is turnable
	targetElem.Color = whoseTurn
	for _, c := range candidate {
		c.Color = whoseTurn
	}

	// Game status update
	b.MyTurn = b.MyTurn * -1
	b.TotalTurn += 1
	return nil
}

func (b *Board) Picture() {
	// Store string as such
	// [Y=1]: X=1, X=2, X=3 ...
	fmt.Printf("[Game turn %v]\n\n", b.TotalTurn)
	pictureRow := map[int][]string{}
	picture := map[int]string{}
	for _, e := range b.Elem {
		repr := " "
		switch e.Color {
		case 1:
			repr = "w"
		case -1:
			repr = "b"
		}
		pictureRow[e.Coordinate.Y] = append(pictureRow[e.Coordinate.Y], repr)
	}

	for k, v := range pictureRow {
		picture[k] = strings.Join(v, "|")
	}

	for i := b.MySize - 1; i >= 0; i-- {
		fmt.Printf("Y > %v :|%s|\n", i, picture[i])
	}

	// Display for bottom column
	rowXNum, rowXEq, rowX := []string{}, []string{}, []string{}
	for i := 0; i < b.MySize; i++ {
		rowXNum = append(rowXNum, fmt.Sprintf("%v", i))
		rowXEq = append(rowXEq, "^")
		rowX = append(rowX, "X")
	}
	fmt.Printf(
		"        %s\n       %s\n       %s\n",
		strings.Join(rowXNum, " "),
		strings.Join(rowXEq, " "),
		strings.Join(rowX, " "),
	)
	fmt.Printf("\nNow Turn: %s\n", StoneColorItoA[b.MyTurn])
	fmt.Printf("White: %v, Black: %v\n", b.Whites, b.Blacks)
}
