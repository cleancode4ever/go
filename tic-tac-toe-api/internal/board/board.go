package board

import (
	"fmt"
)

const dimension = 3

//Board is a 2 dimensional array of CellValues
type Board struct {
	Content [dimension][dimension]
}

//NewBoard Initializes a board with empty cells
func New() Board {
	var content [dimension][dimension]Value

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			content[i][j] = E
		}
	}

	return Board{
		Content: content,
	}
}

//SetCell sets a specific cell of the board to X or O only if it's empty
func (b Board) SetCell(row, col byte, value CellValue) (Board, error) {
	if !isDimensionValid(row) || !isDimensionValid(col) {
		return b, fmt.Errorf("out-of-bound coordination given (%d, %d)", row, col)
	}

	if !isCellValueValid(value) {
		return b, fmt.Errorf("invalid cell value <%d>", value)
	}

	if b.Content[row][col] != E {
		return b, fmt.Errorf("<(%d, %d) = %d> it cannot be set to %d", row, col, b.Content[row][col], value)
	}

	b.Content[row][col] = value

	return b, nil
}

func (b Board) numCellsWithValue(value CellValue) byte {
	var count byte = 0

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if b.Content[i][j] == value {
				count++
			}
		}
	}

	return count
}

func isDimensionValid(dim byte) bool {
	return dim >= 0 && dim <= dimension
}
