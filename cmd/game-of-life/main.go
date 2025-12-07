// Package main
package main

import (
	"fmt"

	"github.com/gentmaks/game-of-life/internal/game"
)

func main() {
	board := game.BoardInit(5, 5)
	it := 0
	for it < 5 {
		fmt.Println("************************************************")
		fmt.Println("Board at iteration: ", it)
		fmt.Println("************************************************")
		board.PrintBoard()
		board.Advance()
		it++
	}
}
