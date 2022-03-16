package domain

import (
	"math/rand"
)

type cell struct {
	coordX int
	coordY int
	number int
}

type Board struct {
	cells [81]cell
}

func (board *Board) defineCoords() {
	currentX, currentY := 1, 1
	for i := range board.cells {
		board.cells[i].coordX = currentX
		board.cells[i].coordY = currentY

		currentX++
		if currentX > 9 {
			currentX = 1
			currentY++
		}
	}
}

func (board *Board) RandomizeBoard() {
	board.defineCoords()

	for i := range board.cells {
		board.cells[i].number = rand.Intn(9) + 1
	}
}
