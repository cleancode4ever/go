package tests

import (
	"github.com/looking-4-justice/tic-tac-toe-api/types"
	"testing"
)

func TestSetCell(t *testing.T) {
	var row, col byte = 0, 1
	var newValue types.CellValue = types.X

	board := types.NewBoard()
	board.Content[row][col].SetCell(newValue)

	value := board.Content[row][col].Cell()

	if value != newValue {
		t.Errorf("Cell(%d, %d) Setter error; %d expected %d received", row, col, newValue, value)
	}
}
