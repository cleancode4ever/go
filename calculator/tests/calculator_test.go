package domain_test

import (
	"fmt"
	"github.com/cleancode4ever/calculator/interfaces"
	"github.com/cleancode4ever/calculator/types"
	"testing"
)

func TestPlainNumberEvaluatesToItself(t *testing.T) {
	const value float64 = 5
	const op = types.Operand(value)
	actual := fmt.Sprintf("%.2f", value)

	evaluatedCalculation := fmt.Sprintf("%.2f", op.Evaluate())

	if evaluatedCalculation != actual {
		t.Errorf("Incorrect calcultion, got: %s, want: %s.", evaluatedCalculation, actual)
	}
}

func TestPlainNumberPrefixInfixPostfix(t *testing.T) {
	const value float64 = 5
	const op = types.Operand(value)
	actual := fmt.Sprintf("%.2f", value)

	if calcPrefix := op.Prefix(); calcPrefix != actual {
		t.Errorf("Wrong prefix notation: %s; expected %s", calcPrefix, actual)
	}

	if calcInfix := op.Infix(); calcInfix != actual {
		t.Errorf("Wrong infix notation: %s; expected %s", calcInfix, actual)
	}

	if calcPostfix := op.Postfix(); calcPostfix != actual {
		t.Errorf("Wrong postfix notation: %s; expected %s", calcPostfix, actual)
	}
}

func TestAddingLiterals(t *testing.T) {
	const left = types.Operand(5)
	const right = types.Operand(15)

	var actual = left.Evaluate() + right.Evaluate()

	evaluatedCalculation := types.NewAdd(left, right).Evaluate()

	if evaluatedCalculation != actual {
		t.Errorf("Incorrect calcultion, got: %.2f, want: %.2f.", evaluatedCalculation, actual)
	}
}

func TestAddingLiteralsPrefixInfixPostfix(t *testing.T) {
	const op byte = '+'
	const leftValue float64 = 5
	const rightValue float64 = 15

	left := types.Operand(leftValue)
	right := types.Operand(rightValue)

	var prefix = fmt.Sprintf("%c %.2f %.2f", op, leftValue, rightValue)
	var infix = fmt.Sprintf("%.2f %c %.2f", leftValue, op, rightValue)
	var postfix = fmt.Sprintf("%.2f %.2f %c", leftValue, rightValue, op)

	expression := types.NewAdd(left, right)

	if calcPrefix := expression.Prefix(); calcPrefix != prefix {
		t.Errorf("Wrong prefix notation: %s; expected %s", calcPrefix, prefix)
	}

	if calcInfix := expression.Infix(); calcInfix != infix {
		t.Errorf("Wrong infix notation: %s; expected %s", calcInfix, infix)
	}

	if calcPostfix := expression.Postfix(); calcPostfix != postfix {
		t.Errorf("Wrong postfix notation: %s; expected %s", calcPostfix, postfix)
	}
}


func TestMuliplyingLiterals(t *testing.T) {
	left := types.Operand(5)
	right := types.Operand(15)

	var actual = left.Evaluate() * right.Evaluate()

	evaluatedCalculation := types.NewMultiply(left, right).Evaluate()

	if evaluatedCalculation != actual {
		t.Errorf("Incorrect calcultion, got: %.2f, want: %.2f.", evaluatedCalculation, actual)
	}
}

func TestMuliplyingLiteralsPrefixInfixPostfix(t *testing.T) {
	const op byte = '*'
	const leftValue float64 = 5
	const rightValue float64 = 15

	left := types.Operand(leftValue)
	right := types.Operand(rightValue)

	var prefix = fmt.Sprintf("%c %.2f %.2f", op, leftValue, rightValue)
	var infix = fmt.Sprintf("%.2f %c %.2f", leftValue, op, rightValue)
	var postfix = fmt.Sprintf("%.2f %.2f %c", leftValue, rightValue, op)

	expression := types.NewMultiply(left, right)

	if calcPrefix := expression.Prefix(); calcPrefix != prefix {
		t.Errorf("Wrong prefix notation: %s; expected %s", calcPrefix, prefix)
	}

	if calcInfix := expression.Infix(); calcInfix != infix {
		t.Errorf("Wrong infix notation: %s; expected %s", calcInfix, infix)
	}

	if calcPostfix := expression.Postfix(); calcPostfix != postfix {
		t.Errorf("Wrong postfix notation: %s; expected %s", calcPostfix, postfix)
	}
}

func TestSimpleAdditionAndMultiplication(t *testing.T) {
	_, _, _, left1, left2, right, expression := simpleAddAndMultiply3Numbers()

	var actual = left1.Evaluate() + left2.Evaluate() * right.Evaluate()

	evaluatedCalculation := expression.Evaluate()

	if evaluatedCalculation != actual {
		t.Errorf("Incorrect calcultion, got: %.2f, want: %.2f.", evaluatedCalculation, actual)
	}
}

func TestSimpleAdditionAndMultiplicationPrefixInfixPostfix(t *testing.T) {
	leftValue1, leftValue2, rightValue, _, _, _, expression := simpleAddAndMultiply3Numbers()

	var prefix = fmt.Sprintf("+ %.2f * %.2f %.2f", leftValue1, leftValue2, rightValue)
	var infix = fmt.Sprintf("%.2f + %.2f * %.2f", leftValue1, leftValue2, rightValue)
	var postfix = fmt.Sprintf("%.2f %.2f %.2f * +", leftValue1, leftValue2, rightValue)

	if calcPrefix := expression.Prefix(); calcPrefix != prefix {
		t.Errorf("Wrong prefix notation: %s; expected %s", calcPrefix, prefix)
	}

	if calcInfix := expression.Infix(); calcInfix != infix {
		t.Errorf("Wrong infix notation: %s; expected %s", calcInfix, infix)
	}

	if calcPostfix := expression.Postfix(); calcPostfix != postfix {
		t.Errorf("Wrong postfix notation: %s; expected %s", calcPostfix, postfix)
	}
}

func TestComplicatedAdditionAndMultiplication(t *testing.T) {
	_, _, _, _, _, A, B, C, D, E, expression := complexAdditionAndMultiplicationOf5Numbers()

	var actual = (A.Evaluate() + B.Evaluate()) * ((C.Evaluate() + D.Evaluate()) * E.Evaluate())

	evaluatedCalculation := expression.Evaluate()

	if evaluatedCalculation != actual {
		t.Errorf("Incorrect calcultion, got: %.2f, want: %.2f.", evaluatedCalculation, actual)
	}
}

func TestComplexAdditionAndMultiplicationPrefixInfixPostfix(t *testing.T) {
	a, b, c, d, e, _, _, _, _, _, expression := complexAdditionAndMultiplicationOf5Numbers()

	var prefix = fmt.Sprintf("* + %.2f %.2f * + %.2f %.2f %.2f", a, b, c, d, e)
	var infix = fmt.Sprintf("%.2f + %.2f * %.2f + %.2f * %.2f", a, b, c, d, e)
	var postfix = fmt.Sprintf("%.2f %.2f + %.2f %.2f + %.2f * *", a, b, c, d, e)

	if calcPrefix := expression.Prefix(); calcPrefix != prefix {
		t.Errorf("Wrong prefix notation: %s; expected %s", calcPrefix, prefix)
	}

	if calcInfix := expression.Infix(); calcInfix != infix {
		t.Errorf("Wrong infix notation: %s; expected %s", calcInfix, infix)
	}

	if calcPostfix := expression.Postfix(); calcPostfix != postfix {
		t.Errorf("Wrong postfix notation: %s; expected %s", calcPostfix, postfix)
	}
}

/*================== Helpers ======================*/

func simpleAddAndMultiply3Numbers() (float64, float64, float64, interfaces.Expression, interfaces.Expression, interfaces.Expression, interfaces.Expression) {
	const leftValue1 float64 = 5
	const leftValue2 float64 = 4
	const rightValue float64 = 15

	left1 := types.Operand(leftValue1)
	left2 := types.Operand(leftValue2)
	right := types.Operand(rightValue)

	expression := types.NewAdd(left1, types.NewMultiply(left2, right))

	return leftValue1, leftValue2, rightValue, left1, left2, right, expression
}

func complexAdditionAndMultiplicationOf5Numbers() (float64, float64, float64, float64, float64, interfaces.Expression, interfaces.Expression, interfaces.Expression, interfaces.Expression, interfaces.Expression, interfaces.Expression) {
	const a float64 = 5
	const b float64 = 4
	const c float64 = 15
	const d float64 = 15
	const e float64 = 15

	A := types.Operand(a)
	B := types.Operand(b)
	C := types.Operand(c)
	D := types.Operand(d)
	E := types.Operand(e)

	expression := types.NewMultiply(types.NewAdd(A, B), types.NewMultiply(types.NewAdd(C, D), E))

	return a, b, c, d, e, A, B, C, D, E, expression
}

