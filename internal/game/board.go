// Package game establishes the structs needed for the game along with the methods for the board
package game

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var directions = []Offset{
	{dx: 0, dy: 1},
	{dx: 1, dy: 0},
	{dx: -1, dy: 0},
	{dx: 0, dy: -1},
	{dx: 1, dy: 1},
	{dx: -1, dy: -1},
	{dx: -1, dy: 1},
	{dx: 1, dy: -1},
}

type Cell struct {
	x     int
	y     int
	state int
}

type Board struct {
	height int
	width  int
	board  [][]Cell
}

type Offset struct {
	dx int
	dy int
}

func BoardInit(x int, y int) *Board {
	board := make([][]Cell, x)
	for i := range x {
		board[i] = make([]Cell, y)
	}
	for row := 0; row < x; row++ {
		for col := 0; col < y; col++ {
			board[row][col] = Cell{x: row, y: col, state: 0}
		}
	}
	for range 1000 {
		randX := rand.IntN(x)
		randY := rand.IntN(y)
		board[randX][randY].state = 1
	}
	return &Board{height: y, width: x, board: board}
}

func (b *Board) GetCell(x int, y int) (Cell, error) {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return Cell{}, errors.New("index out of range for the board")
	}
	return b.board[x][y], nil
}

func (b *Board) SetCell(x int, y int, state int) bool {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return false
	}
	b.board[x][y].state = state
	return true
}

func (b *Board) getNumAliveNeighbors(x int, y int) int {
	res := 0
	for _, d := range directions {
		dx, dy := d.dx, d.dy
		newX, newY := x+dx, y+dy
		if newX < 0 || newX >= b.width || newY < 0 || newY >= b.height {
			continue
		}
		if b.board[newX][newY].state == 1 {
			res++
		}
	}
	return res
}

func (b *Board) Advance() {
	// Create a new board state to avoid modifying while reading
	newState := make([][]int, b.width)
	for i := range b.width {
		newState[i] = make([]int, b.height)
	}

	// Calculate new state for each cell
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			aliveCount := b.getNumAliveNeighbors(i, j)
			if b.board[i][j].state == 1 {
				if aliveCount < 2 || aliveCount > 3 {
					newState[i][j] = 0
				} else {
					newState[i][j] = 1
				}
			} else {
				if aliveCount == 3 {
					newState[i][j] = 1
				} else {
					newState[i][j] = 0
				}
			}
		}
	}

	// Copy new state back to board
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			b.board[i][j].state = newState[i][j]
		}
	}
}

func (b *Board) PrintBoard() {
	for _, row := range b.board {
		for _, val := range row {
			fmt.Printf("%d", val.state)
		}
		fmt.Println("")
	}
}

// GetDimensions returns the width and height of the board
func (b *Board) GetDimensions() (int, int) {
	return b.width, b.height
}

// GetCellState returns the state of a cell at position (x, y)
func (b *Board) GetCellState(x int, y int) int {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return 0
	}
	return b.board[x][y].state
}
