package cell

type Value byte

const (
	O Value = iota
	X
	E
)

func validate(value Value) bool {
	return value == O || value == X || value == E
}

func nextValue(value Value) Value {
	return value ^ X
}
