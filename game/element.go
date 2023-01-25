package game

import (
	"fmt"
	"reflect"
)

var StoneColorItoA map[int]string = map[int]string{
	-1: "Black",
	1:  "White",
	0:  "No stone",
}

var StoneColorAtoI map[string]int = map[string]int{
	"White": 1,
	"Black": -1,
}

type BoardCoord struct {
	X int
	Y int
}

type BoardElement struct {
	// Relative space pieces
	//   |      N
	// Y | W <-- --> E
	//   |_____ S _____
	// 0 |      X
	NW *BoardElement
	N  *BoardElement
	NE *BoardElement
	E  *BoardElement
	W  *BoardElement
	SW *BoardElement
	S  *BoardElement
	SE *BoardElement

	// Space's color -1: Black 0: Empty 1: White
	Color int

	// Board Element's Coordinate
	Coordinate BoardCoord
}

type BoardSpace interface {
	nonEmptyProxy() ([]*BoardElement, error)
	TurnOverCheck(int) ([]*BoardElement, error)
}

func (e *BoardElement) nonEmptyProxy() ([]*BoardElement, error) {
	// Scan 8 directions for Board element.
	// If the board element is not empty return in array of *BoardElement
	myProxy := []*BoardElement{
		e.E, e.W, e.S, e.N,
		e.NE, e.SE, e.NW, e.SW,
	}
	isEmpty := []*BoardElement{}
	for _, p := range myProxy {
		if p.Color == 0 {
			isEmpty = append(isEmpty, p)
		}
	}

	// No more playable space
	if len(isEmpty) == 0 {
		return nil, &GameError{Code: "G999", Message: "Game Over"}
	}

	return isEmpty, nil
}

func traversal(elem *BoardElement, heading string, meet *[]*BoardElement) {
	acrossPtr := reflect.ValueOf(elem)
	across := reflect.Indirect(acrossPtr).FieldByName(heading).Interface().(*BoardElement)
	if across != nil && across.Color != 0 {
		// Add to res if space is not blank
		*meet = append(*meet, across)
		traversal(across, heading, meet)
	} else {
		return
	}
}

func turnOverable(meet []*BoardElement, myColor int) ([]*BoardElement, error) {
	if len(meet) == 0 {
		return nil, &GameError{Code: "GE03", Message: "Cannot turnover anything. No stone present to turnover"}
	}

	for i, space := range meet {
		if space.Color*myColor == 1 {
			// Meet same `myColor` during traversal
			if i == 0 {
				// If it's the first stone I met, there are no stones to turn.
				// Return proxy element error
				return nil, &GameError{Code: "GE01", Message: "Cannot turnover anything, proxy element is same color"}
			}
			// Stones between i'th element and me are turnoverable
			return meet[:i], nil
		}
	}
	return nil, &GameError{Code: "GE02", Message: "Cannot turnover anything, no `myColor` at the end"}
}

func (e *BoardElement) TurnOverCheck(setColor int) ([]*BoardElement, error) {
	heading := []string{"E", "W", "S", "N", "NE", "NW", "SE", "SW"}
	totalTurnOver := []*BoardElement{}
	for _, h := range heading {
		// Check adjacent spaces heading-wise and record cameAcrosses (ca)
		fmt.Println("heading", h)

		ca := []*BoardElement{}
		traversal(e, h, &ca)

		// Among cameAcrosses, select turnover-able (toca)
		toca, err := turnOverable(ca, setColor)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		totalTurnOver = append(totalTurnOver, toca...)
	}
	if len(totalTurnOver) == 0 {
		return nil, &GameError{
			Code:    "GE10",
			Message: fmt.Sprintf("Cannot lay %s stone here", StoneColorItoA[setColor]),
		}
	}

	return totalTurnOver, nil
}
