package domain

import (
	"math"
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
	blocks  [9][9]*cell
}

func (board *Board) defineCoords() {
	currentLine, currentColumn := 1, 1

	for i := range board.cells {
		board.cells[i].line = currentLine
		board.cells[i].column = currentColumn

		board.lines[currentLine-1][currentColumn-1] = &board.cells[i]
		board.columns[currentColumn-1][currentLine-1] = &board.cells[i]
		board.blocks[getBlockByCoord(currentLine, currentColumn)-1][getBlockCellByCoord(currentLine, currentColumn)-1] = &board.cells[i]

		currentLine++
		if currentLine > 9 {
			currentLine = 1
			currentColumn++
		}
	}
}

func getBlockByCoord(line int, column int) int {
	if line <= 3 {
		return int(math.Ceil(float64(column) / 3.0))
	}

	if line <= 6 {
		return int(math.Ceil(float64(column)/3.0)) + 3
	}

	return int(math.Ceil(float64(column)/3.0)) + 6
}

func getBlockCellByCoord(line int, column int) int {
	if line == 1 || line == 4 || line == 7 {
		if column <= 3 {
			return column
		}

		if column <= 6 {
			return column - 3
		}

		return column - 6
	}

	if line == 2 || line == 5 || line == 8 {
		if column <= 3 {
			return column + 3
		}

		if column <= 6 {
			return column
		}

		return column - 3
	}

	if column <= 3 {
		return column + 6
	}

	if column <= 6 {
		return column + 3
	}

	return column
}

func (board *Board) RandomizeBoard() {
	board.defineCoords()

	for i := range board.cells {
		board.cells[i].number = rand.Intn(9) + 1
	}
}
