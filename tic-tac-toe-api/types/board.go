package types

const DIMENSION = 3


type Board struct {
	Dimension byte
	Content [DIMENSION][DIMENSION]Cell
}

func NewBoard() Board {
	return Board{
		Dimension: DIMENSION,
		Content: [DIMENSION][DIMENSION]Cell{
			{NewCell(E), NewCell(E), NewCell(E)},
			{NewCell(E), NewCell(E), NewCell(E)},
			{NewCell(E), NewCell(E), NewCell(E)},
	}}
}