package domain

import (
	"math/rand"
)

type cell struct {
	line   int
	column int
	number int
}

type Board struct {
	cells   [81]cell
	lines   [9][9]*cell
	columns [9][9]*cell
}

func (board *Board) defineCoords() {
	currentLine, currentColumn := 1, 1
	for i := range board.cells {
		board.cells[i].line = currentLine
		board.cells[i].column = currentColumn

		board.lines[currentLine-1][currentColumn-1] = &board.cells[i]
		board.columns[currentColumn-1][currentLine-1] = &board.cells[i]

		currentLine++
		if currentLine > 9 {
			currentLine = 1
			currentColumn++
		}
	}
}

func (board *Board) RandomizeBoard() {
	board.defineCoords()

	for i := range board.cells {
		board.cells[i].number = rand.Intn(9) + 1
	}
}
