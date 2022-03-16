package domain

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
		board.cells[i].number = 1
	}
}
