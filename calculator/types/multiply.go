package types

import "github.com/cleancode4ever/calculator/interfaces"

type Multiply struct {
	BinaryOperator
}

func (bo Multiply) Evaluate() float64 {
	return bo.Left.Evaluate() * bo.Right.Evaluate()
}

func NewMultiply(left interfaces.Expression, right interfaces.Expression) Multiply {
	return Multiply{BinaryOperator{Left: left, Right: right, Operator: '*'}}
}
