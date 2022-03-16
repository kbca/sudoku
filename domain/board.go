package domain

import "math/rand"

type cell struct {
	coordX int
	coordY int
	number int
}

type Board struct {
	cells [81]cell
}

func (board *Board) RandomizeBoard() {
	for i := range board.cells {
		board.cells[i].number = rand.Intn(9) + 1
	}
}
