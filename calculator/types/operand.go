package types

import "fmt"

type Operand float64

func (o Operand) Evaluate() float64 {
	return float64(o)
}

func (o Operand) Prefix() string {
	return fmt.Sprintf("%.2f", o.Evaluate())
}

func (o Operand) Infix() string {
	return fmt.Sprintf("%.2f", o.Evaluate())
}

func (o Operand) Postfix() string {
	return fmt.Sprintf("%.2f", o.Evaluate())
}
