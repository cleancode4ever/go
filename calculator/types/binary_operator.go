package types

import (
	"fmt"
	"github.com/looking-4-justice/calculator/interfaces"
)

type BinaryOperator struct {
	Left     interfaces.Expression
	Operator byte
	Right    interfaces.Expression
}

func (be BinaryOperator) Prefix() string {
	return fmt.Sprintf("%c %s %s", be.Operator, be.Left.Prefix(), be.Right.Prefix())
}

func (be BinaryOperator) Infix() string {
	return fmt.Sprintf("%s %c %s", be.Left.Infix(), be.Operator, be.Right.Infix())
}

func (be BinaryOperator) Postfix() string {
	return fmt.Sprintf("%s %s %c", be.Left.Postfix(), be.Right.Postfix(), be.Operator)
}
