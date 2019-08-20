package interfaces

type Expression interface {
	Prefix() string
	Infix() string
	Postfix() string
	Evaluate() float64
}