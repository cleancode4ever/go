package types

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
