// Package game establishes the structs needed for the game along with the methods for the board
package game

import (
	"errors"
	"fmt"
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
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {
			aliveCount := b.getNumAliveNeighbors(i, j)
			if b.board[i][j].state == 1 {
				if aliveCount < 2 || aliveCount > 3 {
					b.board[i][j].state = 0
				}
			} else {
				if aliveCount == 3 {
					b.board[i][j].state = 1
				}
			}
		}
	}
}

func (b *Board) PrintBoard() {
	fmt.Println(b.board)
}
