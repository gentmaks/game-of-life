// Package main
package main

import (
	"github.com/gentmaks/game-of-life/internal/game"
)

func main() {
	board := game.BoardInit(5, 5)
	for {
		board.PrintBoard()
		board.Advance()
	}
}
