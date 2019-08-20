package interfaces

import "github.com/looking-4-justice/range/types"

type RangerList interface {
	NumberToTitle(number int) ([]string, error)
	TitleToRange(title string) ([]types.IntRange, error)
}