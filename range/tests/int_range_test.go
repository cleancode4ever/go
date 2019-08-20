package int_range_tests

import (
	"fmt"
	"github.com/looking-4-justice/range/types"
	"strings"
	"testing"
)

var intToTitleTestData = []struct {
	number int
	title  string
}{
	{0, "D"},
	{3, "D"},
	{5, "D"},
	{6, "C"},
	{8, "C"},
	{10, "C"},
	{11, "B"},
	{13, "B"},
	{15, "B"},
	{16, "A"},
	{19, "A"},
	{20, "A"},
}

var titleToRangeTestData = []struct{
	title string
	theRange types.IntRange
}{
	{title: "Cold", theRange: types.IntRange{From: -50, To: 0, Title: "Cold"}},
	{title: "Mild", theRange: types.IntRange{From: 0, To: 20, Title: "Mild"}},
	{title: "Warm", theRange: types.IntRange{From: 20, To: 30, Title: "Warm"}},
	{title: "Hot", theRange: types.IntRange{From: 30, To: 50, Title: "Hot"}},
}

func TestExistingNumberToRangeTitleFromToInclusive(t *testing.T) {
	rangeList := getNumberToRangeData(true, true)

	for _, testItem := range intToTitleTestData {
		titles, _ := rangeList.NumberToTitle(testItem.number)
		numTitles := len(titles)

		if numTitles != 1 {
			t.Errorf("Wrong number of titles found: expected 1, received %d", numTitles)
		}

		rangeTitle := titles[0]

		if rangeTitle != testItem.title {
			t.Errorf("Wrong matche found: expected %s, received %s", testItem.title, rangeTitle)
		}
	}
}

func TestTitleToExistingRangeFromToInclusive(t *testing.T) {
	rangeList := getTitleToRangeData(true, true)

	for _, testItem := range titleToRangeTestData {
		ranges, _ := rangeList.TitleToRange(testItem.title)
		numRanges := len(ranges)

		if numRanges != 1 {
			t.Errorf("Wrong number of ranges found: expected 1, received %d", numRanges)
		}

		theRange := ranges[0]

		if theRange != testItem.theRange {
			t.Errorf("Wrong matche found: expected <%v>, received <%v>", testItem.theRange, theRange)
		}
	}
}

func TestNonExistingNumberToRangeTitleFromToInclusiveErrorsOut(t *testing.T) {
	number := 200
	rangeList := getNumberToRangeData(true, true)

	_, err := rangeList.NumberToTitle(number)

	if err == nil {
		t.Errorf("Expected error, none given")
	}

	expectedError := fmt.Sprintf("No match found for %d", number)
	receivedError := err.Error()

	if !strings.Contains(receivedError, expectedError) {
		t.Errorf("Expected error <%s>, received <%s>", expectedError, receivedError)
	}
}

func TestNumberToRangeTitleFromToInclusiveReturnsSliceWhenMoreThanOneMatchFound(t *testing.T) {
	number := 18
	separator := ", "
	expectedJoinedTitles := strings.Join([]string{"A", "E"}, separator)

	rangeList := getNumberToRangeData(true, true)
	rangeList.Ranges = append(rangeList.Ranges, types.IntRange{From: 16, To: 20, Title: "E"})

	matchedTitles, _ := rangeList.NumberToTitle(number)

	joinedTitles := strings.Join(matchedTitles, separator)

	if joinedTitles != expectedJoinedTitles {
		t.Errorf("Wrong matches found: expected %s, received %s", expectedJoinedTitles, joinedTitles)
	}
}

func getNumberToRangeData(fromInclusive bool, toInclusive bool) types.IntRangeList {
	return types.IntRangeList{
		Ranges: []types.IntRange{
			{From: 0, To: 5, Title: "D"},
			{From: 6, To: 10, Title: "C"},
			{From: 11, To: 15, Title: "B"},
			{From: 16, To: 20, Title: "A"},
		},
		FromInclusive: fromInclusive,
		ToInclusive:   toInclusive,
	}
}

func getTitleToRangeData(fromInclusive bool, toInclusive bool) types.IntRangeList {
	return types.IntRangeList{
		Ranges: []types.IntRange{
			{From: -50, To: 0, Title: "Cold"},
			{From: 0, To: 20, Title: "Mild"},
			{From: 20, To: 30, Title: "Warm"},
			{From: 30, To: 50, Title: "Hot"},
		},
		FromInclusive: fromInclusive,
		ToInclusive:   toInclusive,
	}
}
