package types

import (
	"fmt"
	"testing"
)

func TestCanSetEmptyBoardCell(t *testing.T) {
	var row, col byte = 0, 1
	var newValue = X

	initialBoard := NewBoard()
	alteredBoard, err := initialBoard.SetCell(row, col, newValue)

	if err != nil {
		t.Error(err)
	}

	if alteredBoard == initialBoard {
		t.Errorf("Board state didn't change through a valid SET CELL operation: %v", alteredBoard)
	}
}

func TestCannotSetEmptyBoardCellToAnInvalidValue(t *testing.T) {
	var row, col byte = 0, 1
	var newValue CellValue = 0xff //Could be any value other than E (0), O (1), and X (2)

	initialBoard := NewBoard()
	alteredBoard, err := initialBoard.SetCell(row, col, newValue)

	if err == nil {
		t.Error("Setting a cell to an invalid value didn't issue any error")
	}

	if alteredBoard != initialBoard {
		t.Errorf("Board state changed through an invalid operation; expected %v received: %v", initialBoard, alteredBoard)
	}
}

func TestCannotOperateOnOutOfRangeRowAndColumn(t *testing.T) {
	var row, col byte = dimension + 1, dimension + 3
	var newValue = X

	defer func() {
		if recover() != nil {
			t.Error("Setting out-of-bound cell didn't panic")
		}
	}()

	initialBoard := NewBoard()
	alteredBoard, _ := initialBoard.SetCell(row, col, newValue)

	if alteredBoard != initialBoard {
		t.Errorf("Board state changed through an invalid operation; expected %v received: %v", initialBoard, alteredBoard)
	}
}

func TestCannotSetNonEmptyBoardCell(t *testing.T) {
	var row, col byte = 0, 1
	var setValue = X
	var overwriteValue = O

	initialBoard := NewBoard()

	board1, _ := initialBoard.SetCell(row, col, setValue)

	board2, err := board1.SetCell(row, col, overwriteValue)

	if err == nil {
		t.Error(fmt.Sprintf("<(%d, %d) = %d> was overwritten by %d", row, col, setValue, overwriteValue))
	}

	if board1 != board2 {
		t.Errorf("Board state changed through an invalid operation; expected %v received: %v", board1, board2)
	}
}

func TestCanCountBoardCellsWithSpecificValues(t *testing.T) {
	var row1, col1 byte = 0, 1
	var row2, col2 byte = 2, 0
	var expectedNumX byte = 2
	var setValue = O

	board := NewBoard()

	board, _ = board.SetCell(row1, col1, setValue)
	board, _ = board.SetCell(row2, col2, setValue)

	numXs := board.numCellsWithValue(setValue)

	if numXs != expectedNumX {
		t.Errorf("Wrong number of cells returned; expected %d received: %d", expectedNumX, numXs)
	}
}

func TestSetBoardCell(t *testing.T) {
	var row1, col1 byte = 0, 1
	var row2, col2 byte = 2, 0
	var expectedNumX byte = 2
	var setValue = O

	board := NewBoard()

	board, _ = board.SetCell(row1, col1, setValue)
	board, _ = board.SetCell(row2, col2, setValue)

	numXs := board.numCellsWithValue(setValue)

	if numXs != expectedNumX {
		t.Errorf("Wrong number of cells returned; expected %d received: %d", expectedNumX, numXs)
	}
}
