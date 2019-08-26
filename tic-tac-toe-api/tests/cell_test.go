package tests

import (
	"github.com/looking-4-justice/tic-tac-toe-api/types"
	"testing"
)

func TestCanSetEmptyCell(t *testing.T) {
	var row, col byte = 0, 1
	var newValue = types.X

	board := types.NewBoard()
	_ = board.Content[row][col].Set(newValue)

	value := board.Content[row][col].Cell()

	if value != newValue {
		t.Errorf("cell(%d, %d) Setter error; %d expected %d received", row, col, newValue, value)
	}
}

func TestCannotSetNonEmptyCell(t *testing.T) {
	var row, col byte = 0, 1
	var newValue = types.X

	board := types.NewBoard()
	_  = board.Content[row][col].Set(newValue)
	err := board.Content[row][col].Set(types.O)

	if err != nil {
		t.Errorf("Expected error <%s> but did not receive any error", err)
	}

	t.Log(err)
}
