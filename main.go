package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/skkugoon/othelo/game"
)

func main() {
	newBoard := game.Board{}
	newBoard.New(8)
	newBoard.Connect()
	newBoard.Initialize()
	newBoard.Picture()

	iAm := 1

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Enter X and Y pos of your move (Current turn %s): \n", game.StoneColorItoA[iAm])
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
			move, err := processStdIn(strings.Split(text, " "))
			if err != nil {
				fmt.Println(err.Error())
				newBoard.Picture()
				continue
			}
			err = newBoard.Move(move, iAm)
			if err != nil {
				fmt.Println(err.Error())
				newBoard.Picture()
				continue
			}
			// After Move success, print picture
			newBoard.Picture()
			// Next turn
			iAm = iAm * -1
		} else {
			// exit if user entered an empty string
			break
		}
	}
}

func processStdIn(spl []string) (game.BoardCoord, error) {
	c := game.BoardCoord{}
	if len(spl) != 2 {
		return c, &game.GameError{
			Code:    "IO01",
			Message: "Incorrect X, Y placement input",
		}
	}

	res := []int{}
	for _, s := range spl {
		si, err := strconv.Atoi(s)
		if err != nil {
			return c, &game.GameError{
				Code:    "IO02",
				Message: "Wrong X, Y coordinate entered",
			}
		}
		res = append(res, si)
	}

	c.X, c.Y = res[0], res[1]
	return c, nil
}
