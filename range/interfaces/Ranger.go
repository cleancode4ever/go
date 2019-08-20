package interfaces

type IntRanger interface {
	IsInRange(number int, fromInclusive bool, toInclusive bool) bool
}
