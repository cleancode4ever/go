package types

import "github.com/looking-4-justice/calculator/interfaces"

type Add struct {
	BinaryOperator
}

func (bo Add) Evaluate() float64 {
	return bo.Left.Evaluate() + bo.Right.Evaluate()
}

func NewAdd(left interfaces.Expression, right interfaces.Expression) Add {
	return Add{BinaryOperator{Left: left, Right: right, Operator: '+'}}
}