package types

import (
	"errors"
	"fmt"
)

type IntRangeList struct {
	Ranges        []IntRange
	FromInclusive bool
	ToInclusive   bool
}

func (irl IntRangeList) NumberToTitle(number int) ([]string, error) {
	var result []string

	for _, item := range irl.Ranges {
		if item.isInRange(number, irl.FromInclusive, irl.ToInclusive) {
			result = append(result, item.Title)
		}
	}

	if len(result) == 0 {
		return nil, errors.New(fmt.Sprintf("No match found for %d", number))
	}

	return result, nil
}

func (irl IntRangeList) TitleToRange(title string) ([]IntRange, error) {
	var result []IntRange

	for _, item := range irl.Ranges {
		if item.Title == title {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return nil, errors.New(fmt.Sprintf("No match found for %s", title))
	}

	return result, nil
}
