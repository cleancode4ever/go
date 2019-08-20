package types

type IntRange struct {
	From  int
	To    int
	Title string
}

func (r IntRange) isInRange(number int, fromInclusive bool, toInclusive bool) bool {
	if fromInclusive == true && number == r.From {
		return true
	}

	if toInclusive == true && number == r.To {
		return true
	}

	return number > r.From && number < r.To
}