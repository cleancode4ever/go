package game

import "fmt"

type Game struct {
	latestInputCellValue CellValue
	board Board
}

//NewGame Initializes a new game
func NewGame() Game {
	return Game{latestInputCellValue: E, board:NewBoard()}
}

func (g Game) SetBoardCell(row, col byte, value CellValue) (Game, error) {
	if g.latestInputCellValue == value {
		return g, fmt.Errorf("it's not <%d>'s turn", value)
	}

	board, err := g.board.SetCell(row, col, value)

	g.board = board
	g.latestInputCellValue = value

	if g.isWinner(value) {
		return g, fmt.Errorf("<%d> won", value)
	}

	numCellsWithThisValue := g.board.numCellsWithValue(value)
	numCellsWithOtherValue := g.board.numCellsWithValue(NextValidCellValue(value))

	if numCellsWithThisValue == numCellsWithOtherValue && numCellsWithThisValue == dimension {
		return g, fmt.Errorf("draw")
	}

	return g, err
}

func (g Game) isWinner(value CellValue) bool {
	return g.board.CalculateMagicSquareNumberFor(value) == MagicSquareWinnerNumber
}