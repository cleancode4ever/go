package types

type CellValue byte

const (
	E CellValue = iota
	O
	X
)

type Cell struct {
	value CellValue
}

func (c *Cell) SetCell(value CellValue) {
	c.value = value
}

func (c Cell) Cell() CellValue {
	return c.value
}

func NewCell(content CellValue) Cell {
	return Cell{value: content}
}