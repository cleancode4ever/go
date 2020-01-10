package game

import "testing"

func TestOAfterOneXIsAnAcceptableState(t *testing.T) {
	initialGame := NewGame()
	game1, _ := initialGame.SetBoardCell(0, 0, X)
	game2, err := game1.SetBoardCell(0, 1, O)

	if err != nil {
		t.Error(err)
	}

	if initialGame == game1 || initialGame == game2 || game1 == game2 {
		t.Errorf("invalid board state (%v) -> (%v) -> (%v)", initialGame, game1, game2)
	}
}

func TestTwoConsecutiveXIsAnUnacceptableState(t *testing.T) {
	initialGame := NewGame()
	game1, _ := initialGame.SetBoardCell(0, 0, X)
	_, err := game1.SetBoardCell(0, 1, X)

	if err == nil {
		t.Error("Two consecutive Xs should've errored out")
	}
}

func TestCanDetect3XsInARowAsWinner(t *testing.T) {
	initialGame := NewGame()
	game1, _ := initialGame.SetBoardCell(0, 0, X)
	game2, _ := game1.SetBoardCell(2, 1, O)
	game3, _ := game2.SetBoardCell(0, 1, X)
	game4, _ := game3.SetBoardCell(1, 1, O)
	_, err := game4.SetBoardCell(0, 2, X)

	if err == nil {
		t.Error("Win was not detected")
	}
}

func TestCanDetectDraw(t *testing.T) {
	initialGame := NewGame()
	game1, _ := initialGame.SetBoardCell(0, 0, X)
	game2, _ := game1.SetBoardCell(2, 1, O)
	game3, _ := game2.SetBoardCell(0, 1, X)
	game4, _ := game3.SetBoardCell(1, 1, O)
	game5, _ := game4.SetBoardCell(2, 0, X)
	_, err := game5.SetBoardCell(2, 2, O)

	if err == nil {
		t.Error("Draw was not detected")
	}
}
