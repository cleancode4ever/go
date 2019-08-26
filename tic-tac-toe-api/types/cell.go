package types

import "errors"

type CellValue byte

const (
	E CellValue = iota
	O
	X
)

type Cell struct {
	value CellValue
}

func (c Cell) Cell() CellValue {
	return c.value
}

func (c *Cell) Set(value CellValue) error {
	currentVal := c.Cell()

	if currentVal != E {
		return errors.New("Cell is not empty; content: " + string(c.Cell()))
	}

	c.value = value

	return nil
}

func NewCell(value CellValue) Cell {
	return Cell{value: value}
}
